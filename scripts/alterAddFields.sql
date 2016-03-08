select

      CASE
      WHEN (dt.code = "reference")
        THEN
          concat(ea.code, ' ', dt.db_data_type, ' ', dt.addon)
      WHEN (ea.len > 0)
        THEN
          CONCAT(ea.code, ' ', dt.db_data_type, '(', ea.len, ')')
      WHEN (ea.len = 0)
        THEN
          concat(ea.code, ' ', dt.db_data_type, ' ', dt.addon) END ssql

from entity_attrs ea,entities e,data_types dt where e.id=ea.entity_id
                                                    and dt.id=ea.data_type_id
                                                    and e.code='bi_driver_loc'
and not exists
(select 1 from information_schema.columns i where i.table_schema='golang'
  and i.table_name=e.code COLLATE utf8_unicode_ci and i.column_name=ea.code  COLLATE utf8_unicode_ci