ALTER TABLE statuses
DROP COLUMN 'order';

ALTER TABLE statuses
ADD COLUMN 'meta' varchar null;