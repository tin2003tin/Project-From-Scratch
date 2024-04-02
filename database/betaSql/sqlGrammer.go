package betasql

const SQL_GRAMMER string = `
		S' -> S                        
		S -> Select C From T J W        	## loadSql
		S -> Insert Into T CS Value ( V )   ## loadInsert
		S -> Delete From T W 			## deleteRow
		CS -> Columns ( C )				## loadColumns
		CS -> ''
		S -> 
		J -> Join ID On ID Op ID		## loadJoin
		J -> ''                         
		W -> Where Con                  ## loadCondition
		W -> ''  						
		Con -> ID Op Val 				## condition
		V -> ID , V   					## addValue
		V -> ID							## loadValue
		Val -> ID 					## loadType
		C -> ID , C 					## addColumn
		C -> ID 						## loadColumn
		T -> ID 						## loadTable
		Op -> ID 						## loadOp
							`
