import ObjectID from 'bson-objectid'

export const generateObjectId = (): string => {
    return ObjectID().toHexString()
};