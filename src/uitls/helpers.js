const {capitalize} = require('lodash')

function formatEndpoint(path, method, operation){
    return {
        path, 
        method : method.toUpperCase(),
        name : operation.operationId || generateOperationName(path, method),
        description : operation.description,
        params : formatParameters(operation.parameters),
        responses : operation,responses,
    }
}

function generateOperationName(path, method){
    const parts = path.split('/').filter(Boolean)
    const lastPart = parts[parts.length - 1]
    return `${method.toLowerCase()}${capitalize(lastPart)}`;
}

function formatParameters(parameters = []){
    return parameters.reduce((acc, param) => {
        acc[param.name] = param.schema.type;
        return acc
    }, {})
}

module.exports = {
    formatEndpoint,
    generateOperationName,
    formatParameters
}