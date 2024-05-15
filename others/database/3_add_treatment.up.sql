ALTER TABLE `treatment`
ADD `appointment_date` DATETIME DEFAULT NULL
AFTER `detail`;

ALTER TABLE `treatment`
ADD `appointment` varchar(64) DEFAULT NULL
AFTER `appointment_date`;