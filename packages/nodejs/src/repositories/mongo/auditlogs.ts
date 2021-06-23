import { Connection, FilterQuery, QueryOptions } from 'mongoose';
import { storeDefinitions } from '../../stores';
import { domainsAuditLog } from '../../domains';
import { repositoryContainer } from '..';
import {
  getMongoQueryOptions,
  getPagerResults,
  ListWithPager,
} from '../../helpers';

const getModel = (): storeDefinitions.mongo.auditLogs.AuditLogModel => {
  const conn = repositoryContainer.conn as Connection;
  return conn.models
    .auditlogs as storeDefinitions.mongo.auditLogs.AuditLogModel;
};

export const findOneById = async (
  id: string
): Promise<domainsAuditLog.AuditLog | null> => {
  const model = getModel();
  const doc = await model.findById(id);
  return doc ? doc.toJSON() : null;
};

export const find = async (
  conditions: FilterQuery<domainsAuditLog.AuditLog> = {},
  options?: QueryOptions
): Promise<domainsAuditLog.AuditLog[]> => {
  const model = getModel();
  const docs = await model.find(conditions, null, options);
  return docs.map((doc) => doc.toJSON());
};

export const findWithPager = async (
  conditions: FilterQuery<domainsAuditLog.AuditLog> = {},
  size?: number,
  page?: number
): Promise<ListWithPager<domainsAuditLog.AuditLog>> => {
  const options = getMongoQueryOptions(size, page);
  const [list, totalCount] = await Promise.all([
    find(conditions, options),
    count(conditions),
  ]);
  return {
    ...getPagerResults(totalCount, size, page),
    list,
  };
};

export const count = async (
  conditions: FilterQuery<domainsAuditLog.AuditLog> = {}
): Promise<number> => {
  const model = getModel();
  return await model.countDocuments(conditions);
};

export const createOne = async (
  auditLog: domainsAuditLog.AuditLogCreateAttributes
): Promise<domainsAuditLog.AuditLog> => {
  const model = getModel();
  const doc = await model.create(auditLog);
  return doc.toJSON();
};
