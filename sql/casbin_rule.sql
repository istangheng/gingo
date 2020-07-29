BEGIN;
CREATE TABLE IF NOT EXISTS `casbin_rule`
(
    `p_type` varchar(100) DEFAULT NULL,
    `v0`     varchar(100) DEFAULT NULL,
    `v1`     varchar(100) DEFAULT NULL,
    `v2`     varchar(100) DEFAULT NULL,
    `v3`     varchar(100) DEFAULT NULL,
    `v4`     varchar(100) DEFAULT NULL,
    `v5`     varchar(100) DEFAULT NULL
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5)
VALUES ('p', 'Admin', '/ping*', '(GET)|(POST)|(DELETE)|(PUT)', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5)
VALUES ('p', 'Normal', '/ping*', '(GET)|(POST)|(DELETE)|(PUT)', '', '', '');
INSERT INTO casbin_rule (p_type, v0, v1, v2, v3, v4, v5)
VALUES ('p', 'Guest', '/ping*', '(GET)', '', '', '');
COMMIT;