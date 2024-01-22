-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `urls` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(255) NOT NULL,
  `short_url` varchar(255) NOT NULL,
  `hash` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `short_url` (`short_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `urls`;
-- +goose StatementEnd
