import { Document, Model, Schema, SchemaDefinition } from 'mongoose';
import { RevokedToken } from '../../../domains/auth';
import { REVOKED_TOKEN_RETENTION_SEC } from '../../../constants';

export const name = 'revokedtokens';

const schemaDefinition: SchemaDefinition = {
  token: {
    type: Schema.Types.String,
    required: true,
  },
  revokedAt: {
    type: Schema.Types.Date, // for TTL-Index
    index: {
      expires: `${REVOKED_TOKEN_RETENTION_SEC}s`,
    },
  },
  createdAt: {
    type: Schema.Types.Number,
  },
  updatedAt: {
    type: Schema.Types.Number,
  },
};

export interface RevokedTokenDocument extends RevokedToken, Document {
  id: string; // mongoose.Docmentのidがanyなので上書き
}

export type RevokedTokenModel = Model<RevokedTokenDocument>;

export const schema = new Schema<RevokedTokenDocument>(schemaDefinition, {
  autoIndex: true,
  collection: name,
  strict: true,
  timestamps: { currentTime: (): number => Math.floor(Date.now() / 1000) },
  toJSON: { virtuals: true },
  toObject: { virtuals: true },
  id: true,
  versionKey: false,
});
