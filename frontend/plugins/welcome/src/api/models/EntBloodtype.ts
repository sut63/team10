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
    EntBloodtypeEdges,
    EntBloodtypeEdgesFromJSON,
    EntBloodtypeEdgesFromJSONTyped,
    EntBloodtypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBloodtype
 */
export interface EntBloodtype {
    /**
     * Bloodtype holds the value of the "bloodtype" field.
     * @type {string}
     * @memberof EntBloodtype
     */
    bloodtype?: string;
    /**
     * 
     * @type {EntBloodtypeEdges}
     * @memberof EntBloodtype
     */
    edges?: EntBloodtypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBloodtype
     */
    id?: number;
}

export function EntBloodtypeFromJSON(json: any): EntBloodtype {
    return EntBloodtypeFromJSONTyped(json, false);
}

export function EntBloodtypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBloodtype {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bloodtype': !exists(json, 'bloodtype') ? undefined : json['bloodtype'],
        'edges': !exists(json, 'edges') ? undefined : EntBloodtypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBloodtypeToJSON(value?: EntBloodtype | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bloodtype': value.bloodtype,
        'edges': EntBloodtypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}

