-- +goose Up
use pehlione_go;

-- cleanup previous demo records
DELETE FROM product_variants WHERE product_id IN ('prod_001','prod_002','prod_003');
DELETE FROM products WHERE id IN ('prod_001','prod_002','prod_003');

-- seed 30 showcase products (3 per category)
INSERT INTO products (id, name, slug, description, status, created_at, updated_at) VALUES
  ('prod_elec_01', 'Aurora Pulse Earbuds', 'electronics-aurora-pulse', 'Category: Electronics | Ultra-light earbuds tuned for deep focus playlists.', 'active', NOW(3), NOW(3)),
  ('prod_elec_02', 'Lumen Smart Speaker', 'electronics-lumen-speaker', 'Category: Electronics | Minimal speaker with room-aware sound.', 'active', NOW(3), NOW(3)),
  ('prod_elec_03', 'Volt Edge Powerbank', 'electronics-volt-edge', 'Category: Electronics | Slim 20K USB-C power companion.', 'active', NOW(3), NOW(3)),

  ('prod_fash_01', 'Serif Linen Shirt', 'fashion-serif-linen-shirt', 'Category: Fashion | Weightless linen shirt with modern collar.', 'active', NOW(3), NOW(3)),
  ('prod_fash_02', 'Nova Flow Dress', 'fashion-nova-flow-dress', 'Category: Fashion | Midi dress with flowing pleats for day to night.', 'active', NOW(3), NOW(3)),
  ('prod_fash_03', 'Orbit Street Sneaker', 'fashion-orbit-street-sneaker', 'Category: Fashion | Cushion sneaker inspired by retro runners.', 'active', NOW(3), NOW(3)),

  ('prod_home_01', 'Harbor Stone Lamp', 'home-harbor-stone-lamp', 'Category: Home & Living | Sculpted concrete table light.', 'active', NOW(3), NOW(3)),
  ('prod_home_02', 'Calm Nest Blanket', 'home-calm-nest-blanket', 'Category: Home & Living | Oversized knit blanket that keeps warmth.', 'active', NOW(3), NOW(3)),
  ('prod_home_03', 'Mono Arc Wall Art', 'home-mono-arc-wall-art', 'Category: Home & Living | Monochrome metal arc for entryways.', 'active', NOW(3), NOW(3)),

  ('prod_sport_01', 'Momentum Running Tee', 'sports-momentum-running-tee', 'Category: Sports | Seamless tee with mapped ventilation.', 'active', NOW(3), NOW(3)),
  ('prod_sport_02', 'Climb Lite Bottle', 'sports-climb-lite-bottle', 'Category: Sports | Insulated steel bottle with carabiner lid.', 'active', NOW(3), NOW(3)),
  ('prod_sport_03', 'Stride Flex Legging', 'sports-stride-flex-legging', 'Category: Sports | High-rise legging with four-way stretch.', 'active', NOW(3), NOW(3)),

  ('prod_out_01', 'Trail Crest Jacket', 'outdoor-trail-crest-jacket', 'Category: Outdoor | Windproof shell for quick hikes.', 'active', NOW(3), NOW(3)),
  ('prod_out_02', 'Summit Air Hammock', 'outdoor-summit-air-hammock', 'Category: Outdoor | Packable hammock with tree-safe straps.', 'active', NOW(3), NOW(3)),
  ('prod_out_03', 'Cloudstep Camping Mat', 'outdoor-cloudstep-mat', 'Category: Outdoor | Self-inflating mat for elevated sleep.', 'active', NOW(3), NOW(3)),

  ('prod_beauty_01', 'Glowfield Serum', 'beauty-glowfield-serum', 'Category: Beauty | Daily serum with vitamin complex.', 'active', NOW(3), NOW(3)),
  ('prod_beauty_02', 'Oasis Calm Mist', 'beauty-oasis-calm-mist', 'Category: Beauty | Facial mist that hydrates mid-day.', 'active', NOW(3), NOW(3)),
  ('prod_beauty_03', 'Lumen Clay Mask', 'beauty-lumen-clay-mask', 'Category: Beauty | Dual-action clay mask + polish.', 'active', NOW(3), NOW(3)),

  ('prod_kitchen_01', 'Copperline Chef Pan', 'kitchen-copperline-chef-pan', 'Category: Kitchen | Tri-ply pan that sears evenly.', 'active', NOW(3), NOW(3)),
  ('prod_kitchen_02', 'Steamcraft Kettle', 'kitchen-steamcraft-kettle', 'Category: Kitchen | Gooseneck kettle with fast induction base.', 'active', NOW(3), NOW(3)),
  ('prod_kitchen_03', 'Noir Brew Press', 'kitchen-noir-brew-press', 'Category: Kitchen | Matte black press for bold coffee.', 'active', NOW(3), NOW(3)),

  ('prod_office_01', 'Focusloop Desk Mat', 'office-focusloop-mat', 'Category: Office | Vegan leather desk mat with cable channel.', 'active', NOW(3), NOW(3)),
  ('prod_office_02', 'Orbit Note Organizer', 'office-orbit-note-organizer', 'Category: Office | Magnetic module for pens + notes.', 'active', NOW(3), NOW(3)),
  ('prod_office_03', 'Axis Rise Monitor Stand', 'office-axis-rise-stand', 'Category: Office | Maple stand to lift any monitor.', 'active', NOW(3), NOW(3)),

  ('prod_kids_01', 'Pixel Pop Stacker', 'kids-pixel-pop-stacker', 'Category: Kids | Soft stacking toy with contrast colors.', 'active', NOW(3), NOW(3)),
  ('prod_kids_02', 'Comet Ride Balance Bike', 'kids-comet-ride-bike', 'Category: Kids | Balance bike for first adventures.', 'active', NOW(3), NOW(3)),
  ('prod_kids_03', 'Story Glow Lamp', 'kids-story-glow-lamp', 'Category: Kids | Night lamp that projects calm shapes.', 'active', NOW(3), NOW(3)),

  ('prod_acc_01', 'Luxe Loop Belt', 'accessories-luxe-loop-belt', 'Category: Accessories | Full-grain belt with brushed loop.', 'active', NOW(3), NOW(3)),
  ('prod_acc_02', 'Vista Frame Sunglasses', 'accessories-vista-frame-sunglasses', 'Category: Accessories | Acetate frame with UV shield lenses.', 'active', NOW(3), NOW(3)),
  ('prod_acc_03', 'Echo Daypack', 'accessories-echo-daypack', 'Category: Accessories | 18L commuter daypack with padded sleeve.', 'active', NOW(3), NOW(3));

INSERT INTO product_variants (id, product_id, sku, options_json, price_cents, currency, stock, created_at, updated_at) VALUES
  ('var_elec_01', 'prod_elec_01', 'EL-AP-001', '{"color":"Midnight","size":"Std"}', 34900, 'EUR', 120, NOW(3), NOW(3)),
  ('var_elec_02', 'prod_elec_02', 'EL-LS-002', '{"color":"Graphite","size":"Std"}', 18900, 'EUR', 140, NOW(3), NOW(3)),
  ('var_elec_03', 'prod_elec_03', 'EL-VE-003', '{"color":"Charcoal","size":"Std"}', 5900, 'EUR', 210, NOW(3), NOW(3)),

  ('var_fash_01', 'prod_fash_01', 'FA-SH-004', '{"color":"Sand","size":"M"}', 6900, 'EUR', 80, NOW(3), NOW(3)),
  ('var_fash_02', 'prod_fash_02', 'FA-DR-005', '{"color":"Moon","size":"S"}', 8900, 'EUR', 60, NOW(3), NOW(3)),
  ('var_fash_03', 'prod_fash_03', 'FA-SN-006', '{"color":"Slate","size":"42"}', 12900, 'EUR', 110, NOW(3), NOW(3)),

  ('var_home_01', 'prod_home_01', 'HO-LA-007', '{"color":"Stone","size":"Std"}', 7800, 'EUR', 95, NOW(3), NOW(3)),
  ('var_home_02', 'prod_home_02', 'HO-BL-008', '{"color":"Ivory","size":"L"}', 6400, 'EUR', 100, NOW(3), NOW(3)),
  ('var_home_03', 'prod_home_03', 'HO-AR-009', '{"color":"Black","size":"Std"}', 5200, 'EUR', 70, NOW(3), NOW(3)),

  ('var_sport_01', 'prod_sport_01', 'SP-TE-010', '{"color":"Cobalt","size":"M"}', 4500, 'EUR', 130, NOW(3), NOW(3)),
  ('var_sport_02', 'prod_sport_02', 'SP-BO-011', '{"color":"Silver","size":"Std"}', 3200, 'EUR', 150, NOW(3), NOW(3)),
  ('var_sport_03', 'prod_sport_03', 'SP-LE-012', '{"color":"Onyx","size":"S"}', 5500, 'EUR', 90, NOW(3), NOW(3)),

  ('var_out_01', 'prod_out_01', 'OU-JA-013', '{"color":"Forest","size":"M"}', 12900, 'EUR', 75, NOW(3), NOW(3)),
  ('var_out_02', 'prod_out_02', 'OU-HA-014', '{"color":"Sea","size":"Std"}', 9800, 'EUR', 85, NOW(3), NOW(3)),
  ('var_out_03', 'prod_out_03', 'OU-MA-015', '{"color":"Olive","size":"Std"}', 7400, 'EUR', 95, NOW(3), NOW(3)),

  ('var_beauty_01', 'prod_beauty_01', 'BE-SE-016', '{"color":"Clear","size":"30ml"}', 3900, 'EUR', 160, NOW(3), NOW(3)),
  ('var_beauty_02', 'prod_beauty_02', 'BE-MI-017', '{"color":"Mist","size":"50ml"}', 2900, 'EUR', 140, NOW(3), NOW(3)),
  ('var_beauty_03', 'prod_beauty_03', 'BE-MA-018', '{"color":"Clay","size":"Std"}', 3600, 'EUR', 150, NOW(3), NOW(3)),

  ('var_kitchen_01', 'prod_kitchen_01', 'KI-PA-019', '{"color":"Copper","size":"28cm"}', 11500, 'EUR', 70, NOW(3), NOW(3)),
  ('var_kitchen_02', 'prod_kitchen_02', 'KI-KE-020', '{"color":"Steel","size":"Std"}', 8200, 'EUR', 90, NOW(3), NOW(3)),
  ('var_kitchen_03', 'prod_kitchen_03', 'KI-PR-021', '{"color":"Black","size":"Std"}', 6100, 'EUR', 85, NOW(3), NOW(3)),

  ('var_office_01', 'prod_office_01', 'OF-MA-022', '{"color":"Canyon","size":"Std"}', 4800, 'EUR', 120, NOW(3), NOW(3)),
  ('var_office_02', 'prod_office_02', 'OF-OR-023', '{"color":"Frost","size":"Std"}', 3900, 'EUR', 130, NOW(3), NOW(3)),
  ('var_office_03', 'prod_office_03', 'OF-ST-024', '{"color":"Maple","size":"Std"}', 7400, 'EUR', 85, NOW(3), NOW(3)),

  ('var_kids_01', 'prod_kids_01', 'KD-ST-025', '{"color":"Bright","size":"Std"}', 3200, 'EUR', 150, NOW(3), NOW(3)),
  ('var_kids_02', 'prod_kids_02', 'KD-BI-026', '{"color":"Sunrise","size":"Std"}', 8900, 'EUR', 60, NOW(3), NOW(3)),
  ('var_kids_03', 'prod_kids_03', 'KD-LA-027', '{"color":"Glow","size":"Std"}', 3700, 'EUR', 140, NOW(3), NOW(3)),

  ('var_acc_01', 'prod_acc_01', 'AC-BE-028', '{"color":"Walnut","size":"34"}', 4500, 'EUR', 100, NOW(3), NOW(3)),
  ('var_acc_02', 'prod_acc_02', 'AC-SU-029', '{"color":"Amber","size":"Std"}', 6400, 'EUR', 90, NOW(3), NOW(3)),
  ('var_acc_03', 'prod_acc_03', 'AC-BP-030', '{"color":"Navy","size":"18L"}', 8900, 'EUR', 110, NOW(3), NOW(3));

-- +goose Down
use pehlione_go;

DELETE FROM product_variants WHERE id IN (
  'var_elec_01','var_elec_02','var_elec_03',
  'var_fash_01','var_fash_02','var_fash_03',
  'var_home_01','var_home_02','var_home_03',
  'var_sport_01','var_sport_02','var_sport_03',
  'var_out_01','var_out_02','var_out_03',
  'var_beauty_01','var_beauty_02','var_beauty_03',
  'var_kitchen_01','var_kitchen_02','var_kitchen_03',
  'var_office_01','var_office_02','var_office_03',
  'var_kids_01','var_kids_02','var_kids_03',
  'var_acc_01','var_acc_02','var_acc_03'
);

DELETE FROM products WHERE id IN (
  'prod_elec_01','prod_elec_02','prod_elec_03',
  'prod_fash_01','prod_fash_02','prod_fash_03',
  'prod_home_01','prod_home_02','prod_home_03',
  'prod_sport_01','prod_sport_02','prod_sport_03',
  'prod_out_01','prod_out_02','prod_out_03',
  'prod_beauty_01','prod_beauty_02','prod_beauty_03',
  'prod_kitchen_01','prod_kitchen_02','prod_kitchen_03',
  'prod_office_01','prod_office_02','prod_office_03',
  'prod_kids_01','prod_kids_02','prod_kids_03',
  'prod_acc_01','prod_acc_02','prod_acc_03'
);
