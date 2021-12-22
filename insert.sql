INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('order1', 'track1', 'entry', 'ru', '', 'test1','meest1','4',55,'now','2');
INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('order2', 'track2', 'entryy', 'uk', '', 'test2','meest2','5',56,'noww','3');
INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
VALUES ('order3', 'track3', 'entryyy', 'en', '', 'test3','meest3','6',57,'nowww','4');


INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, fk_payments_order)
VALUES ('b563feb7bb4b6tt', '', 'usd', 'wbplay', 2444,4564345,'alpha', 343,34543,0,'order2');
INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, fk_payments_order)
VALUES ('b563feb23427bb4b6tt', '', 'ru', 'wbpay', 243444,456433445,'bank', 43,343,2,'order1');
INSERT INTO payments (transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee, fk_payments_order)
VALUES ('b563feb324237bb4b6tt', '', 'usd', 'wbplay', 23444,4564345,'alpha', 343,34543,0,'order3');


INSERT INTO delivery (name, phone, zip, city, address, region, email, fk_delivery_order)
VALUES ('Test Testov', '9720000000','2639809','Kiryat Mozkin','Ploshad Mira 15','Kraiot','test@gmail.com', 'order1');
INSERT INTO delivery (name, phone, zip, city, address, region, email, fk_delivery_order)
VALUES ('Test Testovich', '9720000000','2639809','Kiryat Mozkin','Ploshad Mira 15','Kraiot','test@gmail.com','order2');
INSERT INTO delivery (name, phone, zip, city, address, region, email, fk_delivery_order)
VALUES ('Test Testov3', '9720000000','2639809','Kiryat Mozkin','Ploshad Mira 15','Kraiot','test@gmail.com','order3');


INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, fk_items_order)
VALUES (34930, 'WBILMTESTTRACK', 43, 'ab4219087a764ae0btest', 'Mascaras',33,'3',22,2222,'Vivienne',202, 'order1');
INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, fk_items_order)
VALUES (349330, 'WBILMTESTTRACK', 43, 'ab4219087a764ae0btest', 'Mascaras',33,'3',22,2222,'Vivienne',202, 'order1');
INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, fk_items_order)
VALUES (32342330, 'WBI222LMTESTTRACK', 43, 'ab4219087a764ae0btest', 'Mascaras2',33,'3',22,2222,'Vivienne',202, 'order2');
INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, fk_items_order)
VALUES (3234320, 'WBILMTE222STTRACK', 43, 'ab4219087a764ae0btest', 'Mascaras2',33,'3',22,2222,'Vivienne',202, 'order3');
INSERT INTO items (chrt_id, track_number, price, rid, name, sale, size, total_price, nm_id, brand, status, fk_items_order)
VALUES (323423430, 'WBILMTES3333TTRACK', 43, 'ab4219087a764ae0btest', 'Mascaras3',33,'3',22,2222,'Vivienne',202, 'order3');
