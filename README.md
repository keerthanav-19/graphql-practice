# graphql-practice

Golang-Graphql server with graphql-go/graphql

Installation : 

	To clone this Repository to your local environment, use the following command:
    
		git clone https://github.com/Keerthanachinnu/graphql-practice.git

To start the app :

	•	Run the command "go run server.go" from the root folder.

Tool used to test the APIS :

	•	Postman

Usage :

	•	Start the application using the command "go run server.go".
	
  •	Execute the GraphQl Query – 
    
		URL :- http://localhost:9091/goengine?query={list{id,name,info,price}} 

Adding New API to the boiler plate :

	•	Write a graphql schema under schemas folder 
                     sample schema – 
					
                    var ProductType = graphql.NewObject(
    					graphql.ObjectConfig{
        					Name: "Product",
        					Fields: graphql.Fields{
            					"id": &graphql.Field{
                					Type: graphql.Int,
            					},
            					"name": &graphql.Field{
                					Type: graphql.String,
            					},
            					"info": &graphql.Field{
                					Type: graphql.String,
            					},
            					"price": &graphql.Field{
                					Type: graphql.Float,
            					},
        					},
    					},
					)


	•	Write a graphql resolver under resolvers folder 
    
		sample resolver – 
        
		func ProductList(params graphql.ResolveParams) (interface{}, error) {
    		return product_data.Products, nil
		}

	•	Register the Query/Mutation under query/mutaion folder respectively 
    
		sample query entry – 
        
		"list": &graphql.Field{
       		Type: graphql.NewList(product_schema.ProductType), //schema
       		Description: "Get product list",
       		Resolve: product_resolvers.ProductList,  //resolver
     	},

Contributing :

	•	Pull requests are welcome. Please make sure to update tests as appropriate.


