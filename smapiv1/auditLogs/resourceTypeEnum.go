package auditlogs

type ResourceTypeEnum string

func ResourceTypeEnum_Skill() ResourceTypeEnum {
	return "Skill"
}
func ResourceTypeEnum_SkillCatalog() ResourceTypeEnum {
	return "SkillCatalog"
}
func ResourceTypeEnum_InSkillProduct() ResourceTypeEnum {
	return "InSkillProduct"
}
func ResourceTypeEnum_Import() ResourceTypeEnum {
	return "Import"
}
func ResourceTypeEnum_Export() ResourceTypeEnum {
	return "Export"
}
