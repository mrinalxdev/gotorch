const axios = require('axios')

class RestParser{
    constructor(config){
        this.config = config;
        this.axiosInstance = axios.create({
            baseURL : config.restConfig.baseURL,
            headers : config.restConfig.timeout || 5000,
        });
    }

    async analyzeEndpoint(endpoint){
        try {
            const response = await this.axiosInstance.request({
                method : endpoint.method,
                url : endpoint.path,
                params : endpoint.sampleParams,
            });

            return this.interSchema(response.data);
        } catch (error) {
            throw new Error(`Failed to analyze endpoint ${endpoint.path} : ${error.message}`)
        }
    }

    interSchema(data){
        if(Array.isArray(data)){
            return {
                type : 'array',
                items : this.interSchema(data[0] || {}),
            };
        }

        if(typeof data === 'object' && data !== null){
            const fields = {};
            for (const [key, value] of Object.entries(data)) {
                fields[key] = this.interSchema(value);
            }
            return { type : 'object', fields};
        }

        return { type : typeof data};
    }
}

module.exports = RestParser