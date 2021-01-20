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
    EntPatientrightstype,
    EntPatientrightstypeFromJSON,
    EntPatientrightstypeFromJSONTyped,
    EntPatientrightstypeToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntAbilitypatientrightsEdges
 */
export interface EntAbilitypatientrightsEdges {
    /**
     * EdgesOfAbilitypatientrightsPatientrightstype holds the value of the EdgesOfAbilitypatientrightsPatientrightstype edge.
     * @type {Array<EntPatientrightstype>}
     * @memberof EntAbilitypatientrightsEdges
     */
    edgesOfAbilitypatientrightsPatientrightstype?: Array<EntPatientrightstype>;
}

export function EntAbilitypatientrightsEdgesFromJSON(json: any): EntAbilitypatientrightsEdges {
    return EntAbilitypatientrightsEdgesFromJSONTyped(json, false);
}

export function EntAbilitypatientrightsEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntAbilitypatientrightsEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edgesOfAbilitypatientrightsPatientrightstype': !exists(json, 'edgesOfAbilitypatientrightsPatientrightstype') ? undefined : ((json['edgesOfAbilitypatientrightsPatientrightstype'] as Array<any>).map(EntPatientrightstypeFromJSON)),
    };
}

export function EntAbilitypatientrightsEdgesToJSON(value?: EntAbilitypatientrightsEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edgesOfAbilitypatientrightsPatientrightstype': value.edgesOfAbilitypatientrightsPatientrightstype === undefined ? undefined : ((value.edgesOfAbilitypatientrightsPatientrightstype as Array<any>).map(EntPatientrightstypeToJSON)),
    };
}


