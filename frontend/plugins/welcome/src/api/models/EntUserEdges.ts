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
        
        'edgesOfDoctor': !exists(json, 'EdgesOfDoctor') ? undefined : EntDoctorFromJSON(json['EdgesOfDoctor']),
        'edgesOfFinancier': !exists(json, 'EdgesOfFinancier') ? undefined : EntFinancierFromJSON(json['EdgesOfFinancier']),
        'edgesOfMedicalrecordstaff': !exists(json, 'EdgesOfMedicalrecordstaff') ? undefined : EntMedicalrecordstaffFromJSON(json['EdgesOfMedicalrecordstaff']),
        'edgesOfNurse': !exists(json, 'EdgesOfNurse') ? undefined : EntNurseFromJSON(json['EdgesOfNurse']),
        'edgesOfUser2registrar': !exists(json, 'EdgesOfUser2registrar') ? undefined : EntRegistrarFromJSON(json['EdgesOfUser2registrar']),
        'edgesOfUserPatientrights': !exists(json, 'EdgesOfUserPatientrights') ? undefined : EntPatientrightsFromJSON(json['EdgesOfUserPatientrights']),
        'edgesOfUserstatus': !exists(json, 'EdgesOfUserstatus') ? undefined : EntUserstatusFromJSON(json['EdgesOfUserstatus']),
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
        
        'EdgesOfDoctor': EntDoctorToJSON(value.edgesOfDoctor),
        'EdgesOfFinancier': EntFinancierToJSON(value.edgesOfFinancier),
        'EdgesOfMedicalrecordstaff': EntMedicalrecordstaffToJSON(value.edgesOfMedicalrecordstaff),
        'EdgesOfNurse': EntNurseToJSON(value.edgesOfNurse),
        'EdgesOfUser2registrar': EntRegistrarToJSON(value.edgesOfUser2registrar),
        'EdgesOfUserPatientrights': EntPatientrightsToJSON(value.edgesOfUserPatientrights),
        'EdgesOfUserstatus': EntUserstatusToJSON(value.edgesOfUserstatus),
    };
}


