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
    EntDoctor,
    EntDoctorFromJSON,
    EntDoctorFromJSONTyped,
    EntDoctorToJSON,
    EntPatientrecord,
    EntPatientrecordFromJSON,
    EntPatientrecordFromJSONTyped,
    EntPatientrecordToJSON,
    EntTypetreatment,
    EntTypetreatmentFromJSON,
    EntTypetreatmentFromJSONTyped,
    EntTypetreatmentToJSON,
    EntUnpaybill,
    EntUnpaybillFromJSON,
    EntUnpaybillFromJSONTyped,
    EntUnpaybillToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTreatmentEdges
 */
export interface EntTreatmentEdges {
    /**
     * 
     * @type {EntDoctor}
     * @memberof EntTreatmentEdges
     */
    edgesOfDoctor?: EntDoctor;
    /**
     * 
     * @type {EntPatientrecord}
     * @memberof EntTreatmentEdges
     */
    edgesOfPatientrecord?: EntPatientrecord;
    /**
     * 
     * @type {EntTypetreatment}
     * @memberof EntTreatmentEdges
     */
    edgesOfTypetreatment?: EntTypetreatment;
    /**
     * 
     * @type {EntUnpaybill}
     * @memberof EntTreatmentEdges
     */
    edgesOfUnpaybills?: EntUnpaybill;
}

export function EntTreatmentEdgesFromJSON(json: any): EntTreatmentEdges {
    return EntTreatmentEdgesFromJSONTyped(json, false);
}

export function EntTreatmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTreatmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edgesOfDoctor': !exists(json, 'EdgesOfDoctor') ? undefined : EntDoctorFromJSON(json['EdgesOfDoctor']),
        'edgesOfPatientrecord': !exists(json, 'EdgesOfPatientrecord') ? undefined : EntPatientrecordFromJSON(json['EdgesOfPatientrecord']),
        'edgesOfTypetreatment': !exists(json, 'EdgesOfTypetreatment') ? undefined : EntTypetreatmentFromJSON(json['EdgesOfTypetreatment']),
        'edgesOfUnpaybills': !exists(json, 'EdgesOfUnpaybills') ? undefined : EntUnpaybillFromJSON(json['EdgesOfUnpaybills']),
    };
}

export function EntTreatmentEdgesToJSON(value?: EntTreatmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'EdgesOfDoctor': EntDoctorToJSON(value.edgesOfDoctor),
        'EdgesOfPatientrecord': EntPatientrecordToJSON(value.edgesOfPatientrecord),
        'EdgesOfTypetreatment': EntTypetreatmentToJSON(value.edgesOfTypetreatment),
        'EdgesOfUnpaybills': EntUnpaybillToJSON(value.edgesOfUnpaybills),
    };
}


