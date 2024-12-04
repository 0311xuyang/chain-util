package dal

import (
	"gorm.io/gorm"
)

// prepare table for test
const (
	multisigAccountSQL   = "CREATE TABLE IF NOT EXISTS `multisig_account` ( `address` varchar(255) NOT NULL, `name` varchar(512) DEFAULT NULL, `threshold` int NOT NULL, `chain_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL, `members` text NOT NULL, `status` varchar(255) NOT NULL, `pubkey` text NOT NULL, `create_time` timestamp NULL DEFAULT NULL, PRIMARY KEY (`address`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	multisigHistorySQL   = "CREATE TABLE IF NOT EXISTS `multisig_history` ( `tx_hash` varchar(255) NOT NULL, `status` int DEFAULT NULL, `create_time` timestamp NOT NULL ON UPDATE CURRENT_TIMESTAMP, `multisig_addr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL, `memo` varchar(256) DEFAULT NULL, `fees` varchar(255) DEFAULT NULL, `height` varchar(255) DEFAULT NULL, `gas_used` varchar(255) DEFAULT NULL, `gas_wanted` varchar(255) DEFAULT NULL, `total_amount` varchar(255) DEFAULT NULL, `sequence_number` varchar(255) DEFAULT NULL, `tx` text, PRIMARY KEY (`tx_hash`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	pendingTxSQL         = "CREATE TABLE IF NOT EXISTS `pending_transaction` ( `id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL, `wallet_address` varchar(255) NOT NULL, `sequence` varchar(255) DEFAULT NULL, `tx_content` text NOT NULL, `signatures` text, `create_time` timestamp NOT NULL, `memo` varchar(256) DEFAULT NULL, `deleted` smallint DEFAULT '0', PRIMARY KEY (`id`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	stakingPercentageSQL = "CREATE TABLE IF NOT EXISTS `staking_percentage` ( `date` date NOT NULL, `point` varchar(10) NOT NULL, `value` bigint DEFAULT NULL, `chain_name` varchar(255) NOT NULL, PRIMARY KEY (`date`,`chain_name`,`point`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	chainMetricsSQL      = "CREATE TABLE IF NOT EXISTS `chain_metrics` ( `unique_delegate` int DEFAULT NULL, `date` date NOT NULL, `chain_name` varchar(255) NOT NULL, `unique_delegation` int DEFAULT NULL, `unbonded_tokens` bigint DEFAULT NULL, `bonded_tokens` bigint DEFAULT NULL, `market_cap` bigint DEFAULT NULL, `total_supply` bigint DEFAULT NULL, `circulating_supply` bigint DEFAULT NULL, `transactions` bigint DEFAULT NULL, `block_time` varchar(255) DEFAULT NULL, `inflation` varchar(255) DEFAULT NULL, `staking_apr` varchar(255) DEFAULT NULL, `txs_per_block` varchar(255) DEFAULT NULL, `community_pool` bigint DEFAULT NULL, `block_height` bigint DEFAULT NULL, `total_stake` bigint DEFAULT NULL, PRIMARY KEY (`chain_name`,`date`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	topDelegationSQL     = "CREATE TABLE IF NOT EXISTS `top_delegation` ( `delegator` varchar(255) NOT NULL, `amount` bigint DEFAULT NULL, `chain_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL, `validator` varchar(255) NOT NULL, PRIMARY KEY (`delegator`,`chain_name`,`validator`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	topUndelegationSQL   = "CREATE TABLE IF NOT EXISTS `top_undelegation` ( `delegator` varchar(255) NOT NULL, `validator` varchar(255) NOT NULL, `amount` bigint DEFAULT NULL, `release_time` timestamp NOT NULL, `chain_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL, PRIMARY KEY (`delegator`,`validator`,`chain_name`) USING BTREE\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
	coinPriceSQL         = "CREATE TABLE IF NOT EXISTS `coin_price` ( `date` date NOT NULL, `price` varchar(255) DEFAULT NULL, `symbol` varchar(255) DEFAULT NULL, `chain_name` varchar(255) NOT NULL, PRIMARY KEY (`date`,`chain_name`)\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"
)

// Prepare prepares the database.
func Prepare(db *gorm.DB) {
	db.Exec(multisigAccountSQL)
	db.Exec(multisigHistorySQL)
	db.Exec(pendingTxSQL)
	db.Exec(stakingPercentageSQL)
	db.Exec(chainMetricsSQL)
	db.Exec(topDelegationSQL)
	db.Exec(topUndelegationSQL)
	db.Exec(coinPriceSQL)
}
