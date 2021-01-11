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
    EntHistorytaking,
    EntHistorytakingFromJSON,
    EntHistorytakingFromJSONTyped,
    EntHistorytakingToJSON,
    EntUser,
    EntUserFromJSON,
    EntUserFromJSONTyped,
    EntUserToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntNurseEdges
 */
export interface EntNurseEdges {
    /**
     * Historytaking holds the value of the historytaking edge.
     * @type {Array<EntHistorytaking>}
     * @memberof EntNurseEdges
     */
    historytaking?: Array<EntHistorytaking>;
    /**
     * 
     * @type {EntUser}
     * @memberof EntNurseEdges
     */
    user?: EntUser;
}

export function EntNurseEdgesFromJSON(json: any): EntNurseEdges {
    return EntNurseEdgesFromJSONTyped(json, false);
}

export function EntNurseEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntNurseEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'historytaking': !exists(json, 'Historytaking') ? undefined : ((json['Historytaking'] as Array<any>).map(EntHistorytakingFromJSON)),
        'user': !exists(json, 'User') ? undefined : EntUserFromJSON(json['User']),
    };
}

export function EntNurseEdgesToJSON(value?: EntNurseEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'historytaking': value.historytaking === undefined ? undefined : ((value.historytaking as Array<any>).map(EntHistorytakingToJSON)),
        'user': EntUserToJSON(value.user),
    };
}


