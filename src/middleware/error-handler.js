class GraphQLError extends Error {
  constructor(message, code, originalError = null) {
    super(message);
    this.code = code;
    this.originalError = originalError;
  }
}

function handleError(error) {
  if (error.response) {
    return new GraphQLError(
      error.response.data.message || "REST API ERROR",
      error.response.this.status,
      error
    );
  }

  if (error.request){
    return new GraphQLError(
        'No resposne from REST API',
        500,
        error
    )
  }

  return new GraphQLError(
    'Internal server error',
    500,
    error,
  );
}

module.exports = {
    GraphQLError,
    handleError
}