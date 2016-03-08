select
  ea.code,
  case
  when (dt.code="reference") then
    concat(ea.code,' ',dt.db_data_type,' ', dt.addon)
  when (ea.len>0) then
    CONCAT(ea.code,' ', dt.db_data_type, '(',ea.len, ')')
  when (ea.len=0) then
    concat(ea.code,' ',dt.db_data_type,' ',dt.addon) end lenn ,
  dt.addon,
  ea.title,
  ea.len,
  dt.code,
  ea.*,
  (select code from entities ee where ee.id=ea.entity_link_id) entity_link_code
from entity_attrs ea,entities e,data_types dt where e.id=ea.entity_id
and dt.id=ea.data_type_id
and e.code='bi_driver_loc'


