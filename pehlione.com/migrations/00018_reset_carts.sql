-- +goose Up
use pehlione_go;
DELETE FROM cart_items;
DELETE FROM carts;

-- +goose Down
use pehlione_go;
DELETE FROM cart_items;
DELETE FROM carts;
