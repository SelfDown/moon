// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNamePubfileAggregationRel = "pubfile_aggregation_rel"

// PubfileAggregationRel mapped from table <pubfile_aggregation_rel>
type PubfileAggregationRel struct {
	AggRelID  string  `gorm:"column:agg_rel_id;primaryKey" json:"agg_rel_id"` // ID
	AggID     *string `gorm:"column:agg_id" json:"agg_id"`                    // ID
	PubfileID *string `gorm:"column:pubfile_id" json:"pubfile_id"`            // ID
	RelType   *string `gorm:"column:rel_type" json:"rel_type"`                // 关系类型（1、部署关系  2回退关系）
}

// TableName PubfileAggregationRel's table name
func (*PubfileAggregationRel) TableName() string {
	return TableNamePubfileAggregationRel
}

func (*PubfileAggregationRel) PrimaryKey() []string {
	return []string{"agg_rel_id"}
}