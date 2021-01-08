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
    EntTreatment,
    EntTreatmentFromJSON,
    EntTreatmentFromJSONTyped,
    EntTreatmentToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntTypetreatmentEdges
 */
export interface EntTypetreatmentEdges {
    /**
     * Treatment holds the value of the treatment edge.
     * @type {Array<EntTreatment>}
     * @memberof EntTypetreatmentEdges
     */
    treatment?: Array<EntTreatment>;
}

export function EntTypetreatmentEdgesFromJSON(json: any): EntTypetreatmentEdges {
    return EntTypetreatmentEdgesFromJSONTyped(json, false);
}

export function EntTypetreatmentEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntTypetreatmentEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'treatment': !exists(json, 'treatment') ? undefined : ((json['treatment'] as Array<any>).map(EntTreatmentFromJSON)),
    };
}

export function EntTypetreatmentEdgesToJSON(value?: EntTypetreatmentEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'treatment': value.treatment === undefined ? undefined : ((value.treatment as Array<any>).map(EntTreatmentToJSON)),
    };
}


