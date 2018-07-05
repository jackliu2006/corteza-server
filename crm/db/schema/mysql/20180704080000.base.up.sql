CREATE TABLE `crm_module` (
 `id` bigint(20) unsigned NOT NULL,
 `name` varchar(64) NOT NULL COMMENT 'The name of the module',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `crm_module_form` (
 `module_id` bigint(20) unsigned NOT NULL,
 `place` tinyint(3) unsigned NOT NULL,
 `type` varchar(64) NOT NULL COMMENT 'The type of the form input field',
 `name` varchar(64) NOT NULL COMMENT 'The name of the field in the form',
 `title` varchar(255) NOT NULL COMMENT 'The label of the form input',
 `placeholder` varchar(255) NOT NULL COMMENT 'Placeholder text (if supported by input)',
 PRIMARY KEY (`module_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8