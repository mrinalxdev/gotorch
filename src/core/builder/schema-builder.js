const { buildSchema } = require('graphql')
const {capitalize} = require('lodash')

class SchemaBuilder {
    constructor(config){
        this.config = config;
    }

    async buildSchema() {}

    generateTypes(){
        return this.config.restConfig.endpoints.map(endpoint => {
            const typeName = this.getTypeName(endpoint);
            const fields = null
        })
    }

    generateTypeFeilds(schema){
        const fields = [];
        
        // this is for the generation of the fields
        // Fetches and throws into a switch case which can determine the type

        for(const [key, value] of Object.entries(schema.fields || {})) {
            const fieldType = this.mapType(value)
            fields.push(`${key}: ${fieldType}`);
        }
        return fields
    }

    mapType(schemaType){
        switch (schemaType.type) {
            case 'string' : return 'String';
            case 'number' : return 'Float';
            case 'integer' : return 'Int';
            case 'boolean' : return 'Boolean';
            case 'array' : return `[${this.mapType(schemaType.items)}]`;
            case 'object' : return capitalize(schemaType.name || 'Object');
            default : return 'String';
        }
    }

    generateQueries(){
        return null
        // queries for the endpoints and others
    }
}