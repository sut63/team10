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
/**
 * 
 * @export
 * @interface ControllersTreatment
 */
export interface ControllersTreatment {
    /**
     * 
     * @type {string}
     * @memberof ControllersTreatment
     */
    datetreat?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersTreatment
     */
    doctorinfo?: number;
    /**
     * 
     * @type {number}
     * @memberof ControllersTreatment
     */
    patientrecord?: number;
    /**
     * 
     * @type {string}
     * @memberof ControllersTreatment
     */
    treatment?: string;
    /**
     * 
     * @type {number}
     * @memberof ControllersTreatment
     */
    typetreatment?: number;
}

export function ControllersTreatmentFromJSON(json: any): ControllersTreatment {
    return ControllersTreatmentFromJSONTyped(json, false);
}

export function ControllersTreatmentFromJSONTyped(json: any, ignoreDiscriminator: boolean): ControllersTreatment {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'datetreat': !exists(json, 'datetreat') ? undefined : json['datetreat'],
        'doctorinfo': !exists(json, 'doctorinfo') ? undefined : json['doctorinfo'],
        'patientrecord': !exists(json, 'patientrecord') ? undefined : json['patientrecord'],
        'treatment': !exists(json, 'treatment') ? undefined : json['treatment'],
        'typetreatment': !exists(json, 'typetreatment') ? undefined : json['typetreatment'],
    };
}

export function ControllersTreatmentToJSON(value?: ControllersTreatment | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'datetreat': value.datetreat,
        'doctorinfo': value.doctorinfo,
        'patientrecord': value.patientrecord,
        'treatment': value.treatment,
        'typetreatment': value.typetreatment,
    };
}

