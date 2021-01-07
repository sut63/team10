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
    EntUserstatusEdges,
    EntUserstatusEdgesFromJSON,
    EntUserstatusEdgesFromJSONTyped,
    EntUserstatusEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserstatus
 */
export interface EntUserstatus {
    /**
     * Userstatus holds the value of the "Userstatus" field.
     * @type {string}
     * @memberof EntUserstatus
     */
    userstatus?: string;
    /**
     * 
     * @type {EntUserstatusEdges}
     * @memberof EntUserstatus
     */
    edges?: EntUserstatusEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntUserstatus
     */
    id?: number;
}

export function EntUserstatusFromJSON(json: any): EntUserstatus {
    return EntUserstatusFromJSONTyped(json, false);
}

export function EntUserstatusFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserstatus {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'userstatus': !exists(json, 'Userstatus') ? undefined : json['Userstatus'],
        'edges': !exists(json, 'edges') ? undefined : EntUserstatusEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntUserstatusToJSON(value?: EntUserstatus | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Userstatus': value.userstatus,
        'edges': EntUserstatusEdgesToJSON(value.edges),
        'id': value.id,
    };
}

