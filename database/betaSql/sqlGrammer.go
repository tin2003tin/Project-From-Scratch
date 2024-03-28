package betasql

const SQL_GRAMMER string = `
		S' -> S                         ## handle1 
		S -> Select C From T J W        ## loadSql
		J -> Join ID On ID Op ID		## loadJoin
		J -> ''                         
		W -> Where Con                  ## loadCondition
		W -> ''  						
		Con -> ID Op Value 				## condition
		Value -> ID 					## loadType
		C -> ID , C 					## addColumn
		C -> ID 						## loadColumn
		T -> ID 						## loadTable
		Op -> ID 						## loadOp
`