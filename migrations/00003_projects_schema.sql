-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS projects(
    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS projects;
-- +goose StatementEnd
