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
    EntAbilitypatientrights,
    EntAbilitypatientrightsFromJSON,
    EntAbilitypatientrightsFromJSONTyped,
    EntAbilitypatientrightsToJSON,
    EntInsurance,
    EntInsuranceFromJSON,
    EntInsuranceFromJSONTyped,
    EntInsuranceToJSON,
    EntMedicalrecordstaff,
    EntMedicalrecordstaffFromJSON,
    EntMedicalrecordstaffFromJSONTyped,
    EntMedicalrecordstaffToJSON,
    EntPatientrecord,
    EntPatientrecordFromJSON,
    EntPatientrecordFromJSONTyped,
    EntPatientrecordToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientrightsEdges
 */
export interface EntPatientrightsEdges {
    /**
     * 
     * @type {EntAbilitypatientrights}
     * @memberof EntPatientrightsEdges
     */
    edgesOfPatientrightsAbilitypatientrights?: EntAbilitypatientrights;
    /**
     * 
     * @type {EntInsurance}
     * @memberof EntPatientrightsEdges
     */
    edgesOfPatientrightsInsurance?: EntInsurance;
    /**
     * 
     * @type {EntMedicalrecordstaff}
     * @memberof EntPatientrightsEdges
     */
    edgesOfPatientrightsMedicalrecordstaff?: EntMedicalrecordstaff;
    /**
     * 
     * @type {EntPatientrecord}
     * @memberof EntPatientrightsEdges
     */
    edgesOfPatientrightsPatientrecord?: EntPatientrecord;
}

export function EntPatientrightsEdgesFromJSON(json: any): EntPatientrightsEdges {
    return EntPatientrightsEdgesFromJSONTyped(json, false);
}

export function EntPatientrightsEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientrightsEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edgesOfPatientrightsAbilitypatientrights': !exists(json, 'EdgesOfPatientrightsAbilitypatientrights') ? undefined : EntAbilitypatientrightsFromJSON(json['EdgesOfPatientrightsAbilitypatientrights']),
        'edgesOfPatientrightsInsurance': !exists(json, 'EdgesOfPatientrightsInsurance') ? undefined : EntInsuranceFromJSON(json['EdgesOfPatientrightsInsurance']),
        'edgesOfPatientrightsMedicalrecordstaff': !exists(json, 'EdgesOfPatientrightsMedicalrecordstaff') ? undefined : EntMedicalrecordstaffFromJSON(json['EdgesOfPatientrightsMedicalrecordstaff']),
        'edgesOfPatientrightsPatientrecord': !exists(json, 'EdgesOfPatientrightsPatientrecord') ? undefined : EntPatientrecordFromJSON(json['EdgesOfPatientrightsPatientrecord']),
    };
}

export function EntPatientrightsEdgesToJSON(value?: EntPatientrightsEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'EdgesOfPatientrightsAbilitypatientrights': EntAbilitypatientrightsToJSON(value.edgesOfPatientrightsAbilitypatientrights),
        'EdgesOfPatientrightsInsurance': EntInsuranceToJSON(value.edgesOfPatientrightsInsurance),
        'EdgesOfPatientrightsMedicalrecordstaff': EntMedicalrecordstaffToJSON(value.edgesOfPatientrightsMedicalrecordstaff),
        'EdgesOfPatientrightsPatientrecord': EntPatientrecordToJSON(value.edgesOfPatientrightsPatientrecord),
    };
}


