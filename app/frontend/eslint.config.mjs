import path from 'node:path'
import { fileURLToPath } from 'node:url'

import { fixupConfigRules } from '@eslint/compat'
import { FlatCompat } from '@eslint/eslintrc'
import js from '@eslint/js'
import tsParser from '@typescript-eslint/parser'
import unusedImports from 'eslint-plugin-unused-imports'
import globals from 'globals'

const __filename = fileURLToPath(import.meta.url)
const __dirname = path.dirname(__filename)
const compat = new FlatCompat({
    baseDirectory: __dirname,
    recommendedConfig: js.configs.recommended,
    allConfig: js.configs.all,
})

export default [
    ...fixupConfigRules(
        compat.extends(
            'plugin:react/recommended',
            'plugin:@typescript-eslint/eslint-recommended',
            'plugin:@typescript-eslint/recommended',
            'plugin:import/recommended',
            'plugin:import/typescript',
            'plugin:react-hooks/recommended',
            'prettier',
            'plugin:prettier/recommended'
        )
    ),
    {
        files: ['**/*.{ts,tsx}'],

        plugins: {
            'unused-imports': unusedImports,
        },

        languageOptions: {
            globals: {
                ...globals.browser,
                ...globals.jest,
            },

            parser: tsParser,
            ecmaVersion: 'latest',
            sourceType: 'module',

            parserOptions: {
                ecmaFeatures: {
                    jsx: true,
                },
            },
        },
        rules: {
            'no-unused-vars': 'off',
            '@typescript-eslint/no-unused-vars': [
                'error',
                {
                    vars: 'all',
                    varsIgnorePattern: '^_|state|action',
                    args: 'after-used',
                    argsIgnorePattern: '^_|state|action',
                },
            ],

            'prettier/prettier': ['error'],

            '@typescript-eslint/explicit-module-boundary-types': 0,
            '@typescript-eslint/no-empty-function': 0,
            '@typescript-eslint/no-explicit-any': 0,
            '@typescript-eslint/no-empty-interface': 0,
            'react-hooks/exhaustive-deps': 0,
            'react/display-name': 0,
            'react/prop-types': ['error'],
            'import/namespace': 0,
            'import/export': 0,
            'react/jsx-curly-brace-presence': ['error', 'never'],
            'import/order': [
                'error',
                {
                    groups: ['builtin', 'external', 'internal'],

                    pathGroups: [
                        {
                            pattern: 'react',
                            group: 'external',
                            position: 'before',
                        },
                    ],

                    pathGroupsExcludedImportTypes: ['react'],
                    'newlines-between': 'always',

                    alphabetize: {
                        order: 'asc',
                        caseInsensitive: true,
                    },
                },
            ],

            'react/jsx-sort-props': [
                'error',
                {
                    callbacksLast: true,
                    shorthandFirst: false,
                    shorthandLast: true,
                    ignoreCase: true,
                    noSortAlphabetically: false,
                    reservedFirst: true,
                },
            ],
        },
    },
]
