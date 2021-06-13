-- Tables creation
CREATE TABLE IF NOT EXISTS items (
    id BIGSERIAL PRIMARY KEY,
    sku TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    price FLOAT NOT NULL,
    qty INTEGER DEFAULT 0
);
CREATE TABLE IF NOT EXISTS promos (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type INTEGER NOT NULL,
    details JSONB
);
CREATE TABLE IF NOT EXISTS item_promo (
    id BIGSERIAL PRIMARY KEY,
    item_id BIGINT NOT NULL REFERENCES items,
    promo_id BIGINT NOT NULL REFERENCES promos
);
-- Seeding items
INSERT INTO items(sku, name, price, qty)
VALUES ('120P90', 'Google Home', 49.99, 10);
INSERT INTO items(sku, name, price, qty)
VALUES ('43N23P', 'MacBook Pro', 5399.99, 5);
INSERT INTO items(sku, name, price, qty)
VALUES ('A304SD', 'Alexa Speaker', 109.5, 10);
INSERT INTO items(sku, name, price, qty)
VALUES ('234234', 'Raspberry Pi B', 30, 2);
-- Seeding promos
INSERT INTO promos(name, type, details)
VALUES ('Bonus Raspberry Pi B', 2, '{"bonus_item_id":"4","bonus_qty":1,"per_qty":1}');
INSERT INTO promos(name, type, details)
VALUES ('Buy 2 Get 1 Google Home', 2, '{"bonus_item_id":"1","bonus_qty":1,"per_qty":2}');
INSERT INTO promos(name, type, details)
VALUES ('Buy 3 Get 10% Off', 1, '{"discount_percentage":0.1,"min_qty":3}');
-- Seeding item_promo for relation
INSERT INTO item_promo(item_id,promo_id)
VALUES (2,1);
INSERT INTO item_promo(item_id,promo_id)
VALUES (1,2);
INSERT INTO item_promo(item_id,promo_id)
VALUES (3,3);