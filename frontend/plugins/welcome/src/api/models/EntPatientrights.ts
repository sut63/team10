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
    EntPatientrightsEdges,
    EntPatientrightsEdgesFromJSON,
    EntPatientrightsEdgesFromJSONTyped,
    EntPatientrightsEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientrights
 */
export interface EntPatientrights {
    /**
     * Permission holds the value of the "Permission" field.
     * @type {string}
     * @memberof EntPatientrights
     */
    permission?: string;
    /**
     * PermissionArea holds the value of the "PermissionArea" field.
     * @type {string}
     * @memberof EntPatientrights
     */
    permissionArea?: string;
    /**
     * PermissionDate holds the value of the "PermissionDate" field.
     * @type {string}
     * @memberof EntPatientrights
     */
    permissionDate?: string;
    /**
     * Responsible holds the value of the "Responsible" field.
     * @type {string}
     * @memberof EntPatientrights
     */
    responsible?: string;
    /**
     * 
     * @type {number}
     * @memberof EntPatientrights
     */
    abilitypatientrightsId?: number;
    /**
     * 
     * @type {EntPatientrightsEdges}
     * @memberof EntPatientrights
     */
    edges?: EntPatientrightsEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntPatientrights
     */
    id?: number;
    /**
     * 
     * @type {number}
     * @memberof EntPatientrights
     */
    insuranceId?: number;
}

export function EntPatientrightsFromJSON(json: any): EntPatientrights {
    return EntPatientrightsFromJSONTyped(json, false);
}

export function EntPatientrightsFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientrights {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'permission': !exists(json, 'Permission') ? undefined : json['Permission'],
        'permissionArea': !exists(json, 'PermissionArea') ? undefined : json['PermissionArea'],
        'permissionDate': !exists(json, 'PermissionDate') ? undefined : json['PermissionDate'],
        'responsible': !exists(json, 'Responsible') ? undefined : json['Responsible'],
        'abilitypatientrightsId': !exists(json, 'abilitypatientrights_id') ? undefined : json['abilitypatientrights_id'],
        'edges': !exists(json, 'edges') ? undefined : EntPatientrightsEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
        'insuranceId': !exists(json, 'insurance_id') ? undefined : json['insurance_id'],
    };
}

export function EntPatientrightsToJSON(value?: EntPatientrights | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'Permission': value.permission,
        'PermissionArea': value.permissionArea,
        'PermissionDate': value.permissionDate,
        'Responsible': value.responsible,
        'abilitypatientrights_id': value.abilitypatientrightsId,
        'edges': EntPatientrightsEdgesToJSON(value.edges),
        'id': value.id,
        'insurance_id': value.insuranceId,
    };
}


