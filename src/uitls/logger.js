const winston = require('winston')

class Logger {
    constructor(){
        this.logger = winston.createLogger({
            level : 'info',
            format : winston.format.combine(
                winston.format.timestamp(),
                winston.format.json()
            ),
            transports : [
                new winston.transports.Console(),
                new winston.transports.File({filename : 'error.log', level : 'error'}),
                new winston.transports.File({ filename : 'combined.log'})
            ]
        });
    }

    info(message, meta={}){
        this.logger.info(message, meta);
    }

    error(message, meta = {}){
        this.logger.warn(message, meta);
    }

    warn(message, meta={}){
        this.logger.warn(message, meta)
    }
}

module.exports = Logger;