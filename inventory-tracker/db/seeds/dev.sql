-- Development-only seed data.
-- Do not use this file for production.

INSERT INTO users (
    id,
    email,
    display_name,
    password_hash,
    role,
    is_active,
    created_at,
    updated_at
)
VALUES (
    '11111111-1111-1111-1111-111111111111',
    'admin@example.com',
    'Development Admin',
    'temporary-dev-hash',
    'ADMIN',
    TRUE,
    NOW(),
    NOW()
)
ON CONFLICT (id) DO NOTHING;


INSERT INTO products (
    id,
    sku,
    name,
    description,
    unit,
    min_stock,
    on_hand,
    is_active,
    version,
    created_by,
    created_at,
    updated_at
)
VALUES
(
    'aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa',
    'PEN-BLK',
    'Black Pen',
    'Black ballpoint pen',
    'PIECE',
    5,
    20,
    TRUE,
    1,
    '11111111-1111-1111-1111-111111111111',
    NOW(),
    NOW()
),
(
    'bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb',
    'PAPER-A4',
    'A4 Paper',
    NULL,
    'PACK',
    10,
    4,
    TRUE,
    1,
    '11111111-1111-1111-1111-111111111111',
    NOW(),
    NOW()
),
(
    'cccccccc-cccc-cccc-cccc-cccccccccccc',
    'TISSUE-01',
    'Tissue Box',
    NULL,
    'BOX',
    5,
    0,
    TRUE,
    1,
    '11111111-1111-1111-1111-111111111111',
    NOW(),
    NOW()
),
(
    'dddddddd-dddd-dddd-dddd-dddddddddddd',
    'OLD-ITEM',
    'Archived Item',
    'Development archived product',
    'PIECE',
    5,
    0,
    FALSE,
    1,
    '11111111-1111-1111-1111-111111111111',
    NOW(),
    NOW()
)
ON CONFLICT (id) DO NOTHING;