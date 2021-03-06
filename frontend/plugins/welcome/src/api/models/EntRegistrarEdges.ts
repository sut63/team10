/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntRegistrarEdges
 */
export interface EntRegistrarEdges {
    /**
     * 
     * @type {EntUser}
     * @memberof EntRegistrarEdges
     */
    edgesOfUser?: EntUser;
}

export function EntRegistrarEdgesFromJSON(json: any): EntRegistrarEdges {
    return EntRegistrarEdgesFromJSONTyped(json, false);
}

export function EntRegistrarEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntRegistrarEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edgesOfUser': !exists(json, 'EdgesOfUser') ? undefined : EntUserFromJSON(json['EdgesOfUser']),
    };
}

export function EntRegistrarEdgesToJSON(value?: EntRegistrarEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'EdgesOfUser': EntUserToJSON(value.edgesOfUser),
    };
}


