-- name: ListProducts :many
SELECT
    id,
    sku,
    name,
    description,
    unit,
    min_stock,
    on_hand,
    is_active,
    CASE
        WHEN on_hand = 0 THEN 'OUT OF STOCK'
        WHEN on_hand < min_stock THEN 'LOW'
        ELSE 'HEALTHY'
    END::text AS stock_status,
    version,
    created_at,
    updated_at
FROM products
WHERE is_active = TRUE
ORDER BY sku ASC, id ASC;

-- name: GetProductByID :one
SELECT
    id,
    sku,
    name,
    description,
    unit,
    min_stock,
    on_hand,
    is_active,
    CASE
        WHEN on_hand = 0 THEN 'OUT OF STOCK'
        WHEN on_hand < min_stock THEN 'LOW'
        ELSE 'HEALTHY'
    END::text AS stock_status,
    version,
    created_at,
    updated_at
FROM products
WHERE id = $1;
