const SwaggerParser = require("swagger-parser")
const {formatEndpoint} = require('../../uitls/helpers')

class SwaggerParser {
    constructor(){
        this.parser = new SwaggerParser();
    }

    async parse(specUrl){
        try {
            const api = await this.parser.parse(specUrl);
            return this.convertToEndpoints(api);
        } catch (error) {
            throw new Error(`Failed to parse Swagger spec : ${error.message}`)
        }
    }

    convertToEndpoints(api){
        const endpoints = [];

        for (const path in api.paths){
            for (const method in api.paths[path]){
                const operation = api.paths[path][method];
                endpoints.push(formatEndpoint(path, method, operation))
            }
        }

        return endpoints
    }
}

module.exports = SwaggerParser