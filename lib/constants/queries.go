package constants

const ListAllSps = `
SELECT proc.specific_schema AS routine_schema,
       proc.specific_name,
       proc.routine_name,
       args.parameter_name,
       args.parameter_mode,
       args.data_type,
       args.udt_name,
       args.parameter_default
FROM information_schema.routines proc
         LEFT JOIN information_schema.parameters args
                   ON proc.specific_schema = args.specific_schema
                       AND proc.specific_name = args.specific_name
		
WHERE proc.routine_schema NOT IN ('pg_catalog', 'information_schema')
  AND proc.routine_type = 'PROCEDURE'
ORDER BY proc.routine_schema,
         proc.routine_name,
         args.ordinal_position;
`

const ListAllCustomTypes = `SELECT CAST(a.attname AS VARCHAR) attr_name, CAST(t.typname AS VARCHAR) type_name, CAST(t.typcategory AS VARCHAR) type_category, CAST(tt.typname AS VARCHAR) attr_type_name, CAST(tt.typcategory AS VARCHAR) attr_type_category
FROM pg_attribute a
    JOIN pg_type t ON a.attrelid = t.typrelid
    JOIN pg_type tt ON a.atttypid = tt.oid
    LEFT JOIN pg_catalog.pg_namespace n ON n.oid = t.typnamespace
WHERE (t.typrelid = 0 OR (SELECT c.relkind = 'c' FROM pg_catalog.pg_class c WHERE c.oid = t.typrelid))
    AND NOT EXISTS(SELECT 1 FROM pg_catalog.pg_type el WHERE el.oid = t.typelem AND el.typarray = t.oid)
    AND n.nspname NOT IN ('pg_catalog', 'information_schema');
`

const ListEnums = `
SELECT pg_type.typname AS enumtype,
     pg_enum.enumlabel AS enumlabel
 FROM pg_type
 JOIN pg_enum
     ON pg_enum.enumtypid = pg_type.oid
     ORDER BY enumtype, enumsortorder
`
