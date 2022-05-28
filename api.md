Bee CLI Spec

------------------------

Items in brackets `[]` are optional.

A `|` char means "or."

Unless otherwise specified, order matters.

------------------------

bee createdb my_db_name ✅

bee deletedb my_db_name ✅

bee updatedb new_db_name

bee listdb

------------------------

bee createtbl my_db_name tablename [column_name:type other_column_name:other_type]
- Allow string, int, float, datetime ?

bee deletetbl my_db_name tablename

bee updatetbl my_db_name old_table_name new_table_name

bee listtbl my_db_name

------------------------

bee createcol my_db_name table_name column_name:type

bee deletecol my_db_name table_name column_name

bee updatecol my_db_name table_name column_name [new_column_name:new_type]

bee listcol my_db_name table_name

------------------------

bee createrow my_db_name table_name [column_name:column_value] [other_column_name:other_column_value]

bee selectrow my_db_name table_name [column_name=is_value] [AND | OR] [other_column_name=is_other_value]

bee updaterow my_db_name table_name column_to_modify:new_value [column_name=is_value] [AND | OR ] [other_column_name=is_other_value]

bee deleterow my_db_name table_name [column_name=is_value] [AND | OR] [other_column_name=is_other_value]

------------------------

Add column restrictions (not null) in later editions...
