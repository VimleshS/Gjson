# Gjson
Package for decoding json

# How to use


	import 	"github.com/VimleshS/GJSON"

Dummy Json.

  	var complexjson = `{
		        "locationId" : "123456789",
		        "locationName" : "Mumbai",
		        "StateName" : "Maharashtra",
		        "PinCode" : "400060",
		        "question" : "What is color of sky",
		        "questionId" : "13425",
		        "answer" : ["BLUE"],
		        "answerId" : [2],
		        "options" : [ 
		            "RED", 
		            "BLUE", 
		            "ORANGE", 
		            "NONE"
		        ]
	}`
	
	
		    
	func main() {
		gjson := GJSON.GjsonDecoder{}
		gjson.PrintJSON([]byte(complexjson))
	}

