// @ts-check
import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';
import angular from 'angular-eslint';
import importX from 'eslint-plugin-import-x';
import { createTypeScriptImportResolver } from 'eslint-import-resolver-typescript';

/**
 * Feature-Sliced layers, ordered low -> high.
 * A layer may only import from layers below it; `from` lists what is off-limits.
 * `core` is excluded: it holds both low-level infra (store, layout, routes) and
 * the composition root (app.config, app.routes), so it spans the whole stack.
 */
const layerZones = [
    {
        target: './src/app/shared',
        from: [
            './src/app/entities',
            './src/app/features',
            './src/app/pages',
            './src/app/modules',
        ],
    },
    {
        target: './src/app/entities',
        from: ['./src/app/features', './src/app/pages', './src/app/modules'],
    },
    {
        target: './src/app/features',
        from: ['./src/app/pages', './src/app/modules'],
    },
    {
        target: './src/app/pages',
        from: ['./src/app/modules'],
    },
];

export default tseslint.config(
    {
        ignores: ['dist/**', '.angular/**', 'coverage/**', 'node_modules/**'],
    },
    {
        files: ['**/*.ts'],
        extends: [
            eslint.configs.recommended,
            ...tseslint.configs.recommended,
            ...angular.configs.tsRecommended,
            // These two presets are load-bearing. Registering the plugin manually
            // via `plugins: { 'import-x': importX }` leaves ExportMap unable to parse
            // child modules, and every graph rule (no-cycle, named, export) then
            // silently reports nothing while appearing to run.
            importX.flatConfigs.recommended,
            importX.flatConfigs.typescript,
        ],
        processor: angular.processInlineTemplates,
        settings: {
            'import-x/resolver-next': [
                createTypeScriptImportResolver({ project: './tsconfig.json' }),
            ],
        },
        rules: {
            // Catches the cycle in-editor, on the file you are editing.
            // allowUnsafeDynamicCyclicDependency mirrors madge's skipAsyncImports:
            // `loadChildren: () => import(...)` is a separate lazy chunk, not an eager edge.
            'import-x/no-cycle': [
                'error',
                {
                    ignoreExternal: true,
                    allowUnsafeDynamicCyclicDependency: true,
                },
            ],

            // Enforces one-way layer flow. This is the rule that catches an
            // entity importing from pages before it ever becomes a cycle.
            'import-x/no-restricted-paths': ['error', { zones: layerZones }],

            // `separate-type-imports` emits `import type { X } from '...'` as its own
            // statement, which erases the runtime edge and lets madge's
            // skipTypeImports actually skip it. Do not switch to inline-type-imports.
            '@typescript-eslint/consistent-type-imports': [
                'error',
                { prefer: 'type-imports', fixStyle: 'separate-type-imports' },
            ],
        },
    },
    {
        files: ['**/*.html'],
        extends: [...angular.configs.templateRecommended],
    },
);
