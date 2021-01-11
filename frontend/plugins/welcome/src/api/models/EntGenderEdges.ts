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
    EntPatientrecord,
    EntPatientrecordFromJSON,
    EntPatientrecordFromJSONTyped,
    EntPatientrecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntGenderEdges
 */
export interface EntGenderEdges {
    /**
     * Patientrecord holds the value of the patientrecord edge.
     * @type {Array<EntPatientrecord>}
     * @memberof EntGenderEdges
     */
    patientrecord?: Array<EntPatientrecord>;
}

export function EntGenderEdgesFromJSON(json: any): EntGenderEdges {
    return EntGenderEdgesFromJSONTyped(json, false);
}

export function EntGenderEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntGenderEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'patientrecord': !exists(json, 'Patientrecord') ? undefined : ((json['Patientrecord'] as Array<any>).map(EntPatientrecordFromJSON)),
    };
}

export function EntGenderEdgesToJSON(value?: EntGenderEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'patientrecord': value.patientrecord === undefined ? undefined : ((value.patientrecord as Array<any>).map(EntPatientrecordToJSON)),
    };
}


