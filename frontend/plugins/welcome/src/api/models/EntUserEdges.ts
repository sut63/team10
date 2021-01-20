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
    EntFinancier,
    EntFinancierFromJSON,
    EntFinancierFromJSONTyped,
    EntFinancierToJSON,
    EntMedicalrecordstaff,
    EntMedicalrecordstaffFromJSON,
    EntMedicalrecordstaffFromJSONTyped,
    EntMedicalrecordstaffToJSON,
    EntNurse,
    EntNurseFromJSON,
    EntNurseFromJSONTyped,
    EntNurseToJSON,
    EntPatientrights,
    EntPatientrightsFromJSON,
    EntPatientrightsFromJSONTyped,
    EntPatientrightsToJSON,
    EntRegistrar,
    EntRegistrarFromJSON,
    EntRegistrarFromJSONTyped,
    EntRegistrarToJSON,
    EntUserstatus,
    EntUserstatusFromJSON,
    EntUserstatusFromJSONTyped,
    EntUserstatusToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntUserEdges
 */
export interface EntUserEdges {
    /**
     * 
     * @type {EntDoctor}
     * @memberof EntUserEdges
     */
    edgesOfDoctor?: EntDoctor;
    /**
     * 
     * @type {EntFinancier}
     * @memberof EntUserEdges
     */
    edgesOfFinancier?: EntFinancier;
    /**
     * 
     * @type {EntMedicalrecordstaff}
     * @memberof EntUserEdges
     */
    edgesOfMedicalrecordstaff?: EntMedicalrecordstaff;
    /**
     * 
     * @type {EntNurse}
     * @memberof EntUserEdges
     */
    edgesOfNurse?: EntNurse;
    /**
     * 
     * @type {EntRegistrar}
     * @memberof EntUserEdges
     */
    edgesOfUser2registrar?: EntRegistrar;
    /**
     * 
     * @type {EntPatientrights}
     * @memberof EntUserEdges
     */
    edgesOfUserPatientrights?: EntPatientrights;
    /**
     * 
     * @type {EntUserstatus}
     * @memberof EntUserEdges
     */
    edgesOfUserstatus?: EntUserstatus;
}

export function EntUserEdgesFromJSON(json: any): EntUserEdges {
    return EntUserEdgesFromJSONTyped(json, false);
}

export function EntUserEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntUserEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'edgesOfDoctor': !exists(json, 'edgesOfDoctor') ? undefined : EntDoctorFromJSON(json['edgesOfDoctor']),
        'edgesOfFinancier': !exists(json, 'edgesOfFinancier') ? undefined : EntFinancierFromJSON(json['edgesOfFinancier']),
        'edgesOfMedicalrecordstaff': !exists(json, 'edgesOfMedicalrecordstaff') ? undefined : EntMedicalrecordstaffFromJSON(json['edgesOfMedicalrecordstaff']),
        'edgesOfNurse': !exists(json, 'edgesOfNurse') ? undefined : EntNurseFromJSON(json['edgesOfNurse']),
        'edgesOfUser2registrar': !exists(json, 'edgesOfUser2registrar') ? undefined : EntRegistrarFromJSON(json['edgesOfUser2registrar']),
        'edgesOfUserPatientrights': !exists(json, 'edgesOfUserPatientrights') ? undefined : EntPatientrightsFromJSON(json['edgesOfUserPatientrights']),
        'edgesOfUserstatus': !exists(json, 'edgesOfUserstatus') ? undefined : EntUserstatusFromJSON(json['edgesOfUserstatus']),
    };
}

export function EntUserEdgesToJSON(value?: EntUserEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'edgesOfDoctor': EntDoctorToJSON(value.edgesOfDoctor),
        'edgesOfFinancier': EntFinancierToJSON(value.edgesOfFinancier),
        'edgesOfMedicalrecordstaff': EntMedicalrecordstaffToJSON(value.edgesOfMedicalrecordstaff),
        'edgesOfNurse': EntNurseToJSON(value.edgesOfNurse),
        'edgesOfUser2registrar': EntRegistrarToJSON(value.edgesOfUser2registrar),
        'edgesOfUserPatientrights': EntPatientrightsToJSON(value.edgesOfUserPatientrights),
        'edgesOfUserstatus': EntUserstatusToJSON(value.edgesOfUserstatus),
    };
}


