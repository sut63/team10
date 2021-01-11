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
    doctor?: EntDoctor;
    /**
     * 
     * @type {EntPatientrecord}
     * @memberof EntTreatmentEdges
     */
    patientrecord?: EntPatientrecord;
    /**
     * 
     * @type {EntTypetreatment}
     * @memberof EntTreatmentEdges
     */
    typetreatment?: EntTypetreatment;
    /**
     * 
     * @type {EntUnpaybill}
     * @memberof EntTreatmentEdges
     */
    unpaybills?: EntUnpaybill;
}

export function EntTreatmentEdgesFromJSON(json: any): EntTreatmentEdges {
    return EntTreatmentEdgesFromJSONTyped(json, false);
}

export function EntTreatmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTreatmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'doctor': !exists(json, 'Doctor') ? undefined : EntDoctorFromJSON(json['Doctor']),
        'patientrecord': !exists(json, 'Patientrecord') ? undefined : EntPatientrecordFromJSON(json['Patientrecord']),
        'typetreatment': !exists(json, 'Typetreatment') ? undefined : EntTypetreatmentFromJSON(json['Typetreatment']),
        'unpaybills': !exists(json, 'Unpaybills') ? undefined : EntUnpaybillFromJSON(json['Unpaybills']),
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
        
        'doctor': EntDoctorToJSON(value.doctor),
        'patientrecord': EntPatientrecordToJSON(value.patientrecord),
        'typetreatment': EntTypetreatmentToJSON(value.typetreatment),
        'unpaybills': EntUnpaybillToJSON(value.unpaybills),
    };
}


