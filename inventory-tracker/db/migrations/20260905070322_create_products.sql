-- +goose Up

CREATE TABLE products (
    id UUID PRIMARY KEY,
    sku VARCHAR(32) NOT NULL UNIQUE,
    name VARCHAR(100) NOT NULL,
    description VARCHAR(500),
    unit VARCHAR(10) NOT NULL,
    min_stock INTEGER NOT NULL DEFAULT 0,
    on_hand INTEGER NOT NULL DEFAULT 0,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    version INTEGER NOT NULL DEFAULT 1,

    created_by UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT products_created_by_fk
        FOREIGN KEY (created_by)
        REFERENCES users(id)
        ON DELETE RESTRICT,

    CONSTRAINT products_unit_check
        CHECK (unit IN ('PIECE', 'BOX', 'PACK')),

    CONSTRAINT products_min_stock_check
        CHECK (min_stock BETWEEN 0 AND 1000000),

    CONSTRAINT products_on_hand_check
        CHECK (on_hand BETWEEN 0 AND 1000000),

    CONSTRAINT products_version_check
        CHECK (version > 0)
);

CREATE INDEX idx_products_active_sku_id
    ON products (is_active, sku, id);

-- +goose Down

DROP TABLE products;