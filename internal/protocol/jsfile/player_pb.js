/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

goog.exportSymbol('proto.pb.PlayerInfo', null, global);
goog.exportSymbol('proto.pb.PlayerListInfo', null, global);
goog.exportSymbol('proto.pb.PlayerRecord', null, global);
goog.exportSymbol('proto.pb.PlayerState', null, global);
goog.exportSymbol('proto.pb.UpdateMoneyReq', null, global);
goog.exportSymbol('proto.pb.UpdateMoneyResp', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.PlayerInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.PlayerInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.PlayerInfo.displayName = 'proto.pb.PlayerInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.PlayerListInfo = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.pb.PlayerListInfo.repeatedFields_, null);
};
goog.inherits(proto.pb.PlayerListInfo, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.PlayerListInfo.displayName = 'proto.pb.PlayerListInfo';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.PlayerRecord = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.PlayerRecord, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.PlayerRecord.displayName = 'proto.pb.PlayerRecord';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.UpdateMoneyReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.UpdateMoneyReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.UpdateMoneyReq.displayName = 'proto.pb.UpdateMoneyReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.pb.UpdateMoneyResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.pb.UpdateMoneyResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.pb.UpdateMoneyResp.displayName = 'proto.pb.UpdateMoneyResp';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.PlayerInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.PlayerInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.PlayerInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PlayerInfo.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    account: jspb.Message.getFieldWithDefault(msg, 2, ""),
    name: jspb.Message.getFieldWithDefault(msg, 3, ""),
    faceid: jspb.Message.getFieldWithDefault(msg, 4, 0),
    age: jspb.Message.getFieldWithDefault(msg, 5, 0),
    sex: jspb.Message.getFieldWithDefault(msg, 6, 0),
    yuanbao: jspb.Message.getFieldWithDefault(msg, 7, 0),
    coin: jspb.Message.getFieldWithDefault(msg, 8, 0),
    level: jspb.Message.getFieldWithDefault(msg, 9, 0),
    ranking: jspb.Message.getFieldWithDefault(msg, 10, 0),
    state: jspb.Message.getFieldWithDefault(msg, 11, 0),
    gold: jspb.Message.getFieldWithDefault(msg, 12, 0),
    money: jspb.Message.getFieldWithDefault(msg, 13, 0),
    bindinfo: jspb.Message.getFieldWithDefault(msg, 14, ""),
    gamestate: jspb.Message.getFieldWithDefault(msg, 15, 0),
    platformid: jspb.Message.getFieldWithDefault(msg, 16, 0),
    roomnum: jspb.Message.getFieldWithDefault(msg, 17, 0),
    gameid: jspb.Message.getFieldWithDefault(msg, 18, 0),
    tableid: jspb.Message.getFieldWithDefault(msg, 19, 0),
    chairid: jspb.Message.getFieldWithDefault(msg, 20, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.PlayerInfo}
 */
proto.pb.PlayerInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.PlayerInfo;
  return proto.pb.PlayerInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.PlayerInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.PlayerInfo}
 */
proto.pb.PlayerInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setAccount(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setFaceid(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setAge(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setSex(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setYuanbao(value);
      break;
    case 8:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setCoin(value);
      break;
    case 9:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setLevel(value);
      break;
    case 10:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRanking(value);
      break;
    case 11:
      var value = /** @type {!proto.pb.PlayerState} */ (reader.readEnum());
      msg.setState(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGold(value);
      break;
    case 13:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setMoney(value);
      break;
    case 14:
      var value = /** @type {string} */ (reader.readString());
      msg.setBindinfo(value);
      break;
    case 15:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setGamestate(value);
      break;
    case 16:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setPlatformid(value);
      break;
    case 17:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setRoomnum(value);
      break;
    case 18:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setGameid(value);
      break;
    case 19:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTableid(value);
      break;
    case 20:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setChairid(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.PlayerInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.PlayerInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.PlayerInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PlayerInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getAccount();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getFaceid();
  if (f !== 0) {
    writer.writeInt32(
      4,
      f
    );
  }
  f = message.getAge();
  if (f !== 0) {
    writer.writeInt32(
      5,
      f
    );
  }
  f = message.getSex();
  if (f !== 0) {
    writer.writeInt32(
      6,
      f
    );
  }
  f = message.getYuanbao();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getCoin();
  if (f !== 0) {
    writer.writeInt64(
      8,
      f
    );
  }
  f = message.getLevel();
  if (f !== 0) {
    writer.writeInt32(
      9,
      f
    );
  }
  f = message.getRanking();
  if (f !== 0) {
    writer.writeInt32(
      10,
      f
    );
  }
  f = message.getState();
  if (f !== 0.0) {
    writer.writeEnum(
      11,
      f
    );
  }
  f = message.getGold();
  if (f !== 0) {
    writer.writeInt64(
      12,
      f
    );
  }
  f = message.getMoney();
  if (f !== 0) {
    writer.writeInt64(
      13,
      f
    );
  }
  f = message.getBindinfo();
  if (f.length > 0) {
    writer.writeString(
      14,
      f
    );
  }
  f = message.getGamestate();
  if (f !== 0) {
    writer.writeInt32(
      15,
      f
    );
  }
  f = message.getPlatformid();
  if (f !== 0) {
    writer.writeInt64(
      16,
      f
    );
  }
  f = message.getRoomnum();
  if (f !== 0) {
    writer.writeInt64(
      17,
      f
    );
  }
  f = message.getGameid();
  if (f !== 0) {
    writer.writeInt64(
      18,
      f
    );
  }
  f = message.getTableid();
  if (f !== 0) {
    writer.writeInt32(
      19,
      f
    );
  }
  f = message.getChairid();
  if (f !== 0) {
    writer.writeInt32(
      20,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional string account = 2;
 * @return {string}
 */
proto.pb.PlayerInfo.prototype.getAccount = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.pb.PlayerInfo.prototype.setAccount = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string name = 3;
 * @return {string}
 */
proto.pb.PlayerInfo.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.pb.PlayerInfo.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional int32 faceID = 4;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getFaceid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setFaceid = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int32 age = 5;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getAge = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setAge = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int32 sex = 6;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getSex = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setSex = function(value) {
  jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int64 yuanBao = 7;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getYuanbao = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setYuanbao = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * optional int64 coin = 8;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getCoin = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 8, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setCoin = function(value) {
  jspb.Message.setProto3IntField(this, 8, value);
};


/**
 * optional int32 level = 9;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getLevel = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 9, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setLevel = function(value) {
  jspb.Message.setProto3IntField(this, 9, value);
};


/**
 * optional int32 ranking = 10;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getRanking = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 10, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setRanking = function(value) {
  jspb.Message.setProto3IntField(this, 10, value);
};


/**
 * optional PlayerState state = 11;
 * @return {!proto.pb.PlayerState}
 */
proto.pb.PlayerInfo.prototype.getState = function() {
  return /** @type {!proto.pb.PlayerState} */ (jspb.Message.getFieldWithDefault(this, 11, 0));
};


/** @param {!proto.pb.PlayerState} value */
proto.pb.PlayerInfo.prototype.setState = function(value) {
  jspb.Message.setProto3EnumField(this, 11, value);
};


/**
 * optional int64 gold = 12;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getGold = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setGold = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};


/**
 * optional int64 money = 13;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getMoney = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 13, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setMoney = function(value) {
  jspb.Message.setProto3IntField(this, 13, value);
};


/**
 * optional string bindInfo = 14;
 * @return {string}
 */
proto.pb.PlayerInfo.prototype.getBindinfo = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 14, ""));
};


/** @param {string} value */
proto.pb.PlayerInfo.prototype.setBindinfo = function(value) {
  jspb.Message.setProto3StringField(this, 14, value);
};


/**
 * optional int32 gameState = 15;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getGamestate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 15, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setGamestate = function(value) {
  jspb.Message.setProto3IntField(this, 15, value);
};


/**
 * optional int64 platformID = 16;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getPlatformid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 16, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setPlatformid = function(value) {
  jspb.Message.setProto3IntField(this, 16, value);
};


/**
 * optional int64 roomNum = 17;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getRoomnum = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 17, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setRoomnum = function(value) {
  jspb.Message.setProto3IntField(this, 17, value);
};


/**
 * optional int64 gameID = 18;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getGameid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 18, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setGameid = function(value) {
  jspb.Message.setProto3IntField(this, 18, value);
};


/**
 * optional int32 tableID = 19;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getTableid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 19, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setTableid = function(value) {
  jspb.Message.setProto3IntField(this, 19, value);
};


/**
 * optional int32 chairID = 20;
 * @return {number}
 */
proto.pb.PlayerInfo.prototype.getChairid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 20, 0));
};


/** @param {number} value */
proto.pb.PlayerInfo.prototype.setChairid = function(value) {
  jspb.Message.setProto3IntField(this, 20, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.pb.PlayerListInfo.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.PlayerListInfo.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.PlayerListInfo.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.PlayerListInfo} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PlayerListInfo.toObject = function(includeInstance, msg) {
  var obj = {
    allinfosList: jspb.Message.toObjectList(msg.getAllinfosList(),
    proto.pb.PlayerInfo.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.PlayerListInfo}
 */
proto.pb.PlayerListInfo.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.PlayerListInfo;
  return proto.pb.PlayerListInfo.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.PlayerListInfo} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.PlayerListInfo}
 */
proto.pb.PlayerListInfo.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.PlayerInfo;
      reader.readMessage(value,proto.pb.PlayerInfo.deserializeBinaryFromReader);
      msg.addAllinfos(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.PlayerListInfo.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.PlayerListInfo.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.PlayerListInfo} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PlayerListInfo.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAllinfosList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.pb.PlayerInfo.serializeBinaryToWriter
    );
  }
};


/**
 * repeated PlayerInfo allInfos = 1;
 * @return {!Array<!proto.pb.PlayerInfo>}
 */
proto.pb.PlayerListInfo.prototype.getAllinfosList = function() {
  return /** @type{!Array<!proto.pb.PlayerInfo>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.pb.PlayerInfo, 1));
};


/** @param {!Array<!proto.pb.PlayerInfo>} value */
proto.pb.PlayerListInfo.prototype.setAllinfosList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.pb.PlayerInfo=} opt_value
 * @param {number=} opt_index
 * @return {!proto.pb.PlayerInfo}
 */
proto.pb.PlayerListInfo.prototype.addAllinfos = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.pb.PlayerInfo, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 */
proto.pb.PlayerListInfo.prototype.clearAllinfosList = function() {
  this.setAllinfosList([]);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.PlayerRecord.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.PlayerRecord.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.PlayerRecord} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PlayerRecord.toObject = function(includeInstance, msg) {
  var obj = {
    user: (f = msg.getUser()) && proto.pb.PlayerInfo.toObject(includeInstance, f),
    twice: jspb.Message.getFieldWithDefault(msg, 2, 0),
    ranking: jspb.Message.getFieldWithDefault(msg, 3, 0),
    bankroll: jspb.Message.getFieldWithDefault(msg, 4, 0),
    winlos: jspb.Message.getFieldWithDefault(msg, 5, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.PlayerRecord}
 */
proto.pb.PlayerRecord.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.PlayerRecord;
  return proto.pb.PlayerRecord.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.PlayerRecord} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.PlayerRecord}
 */
proto.pb.PlayerRecord.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.pb.PlayerInfo;
      reader.readMessage(value,proto.pb.PlayerInfo.deserializeBinaryFromReader);
      msg.setUser(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setTwice(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRanking(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setBankroll(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setWinlos(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.PlayerRecord.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.PlayerRecord.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.PlayerRecord} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.PlayerRecord.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUser();
  if (f != null) {
    writer.writeMessage(
      1,
      f,
      proto.pb.PlayerInfo.serializeBinaryToWriter
    );
  }
  f = message.getTwice();
  if (f !== 0) {
    writer.writeInt32(
      2,
      f
    );
  }
  f = message.getRanking();
  if (f !== 0) {
    writer.writeInt32(
      3,
      f
    );
  }
  f = message.getBankroll();
  if (f !== 0) {
    writer.writeInt64(
      4,
      f
    );
  }
  f = message.getWinlos();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
};


/**
 * optional PlayerInfo user = 1;
 * @return {?proto.pb.PlayerInfo}
 */
proto.pb.PlayerRecord.prototype.getUser = function() {
  return /** @type{?proto.pb.PlayerInfo} */ (
    jspb.Message.getWrapperField(this, proto.pb.PlayerInfo, 1));
};


/** @param {?proto.pb.PlayerInfo|undefined} value */
proto.pb.PlayerRecord.prototype.setUser = function(value) {
  jspb.Message.setWrapperField(this, 1, value);
};


/**
 * Clears the message field making it undefined.
 */
proto.pb.PlayerRecord.prototype.clearUser = function() {
  this.setUser(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.pb.PlayerRecord.prototype.hasUser = function() {
  return jspb.Message.getField(this, 1) != null;
};


/**
 * optional int32 twice = 2;
 * @return {number}
 */
proto.pb.PlayerRecord.prototype.getTwice = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.PlayerRecord.prototype.setTwice = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int32 ranking = 3;
 * @return {number}
 */
proto.pb.PlayerRecord.prototype.getRanking = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.pb.PlayerRecord.prototype.setRanking = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional int64 bankroll = 4;
 * @return {number}
 */
proto.pb.PlayerRecord.prototype.getBankroll = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.pb.PlayerRecord.prototype.setBankroll = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional int64 winLos = 5;
 * @return {number}
 */
proto.pb.PlayerRecord.prototype.getWinlos = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.pb.PlayerRecord.prototype.setWinlos = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.UpdateMoneyReq.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.UpdateMoneyReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.UpdateMoneyReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.UpdateMoneyReq.toObject = function(includeInstance, msg) {
  var obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.UpdateMoneyReq}
 */
proto.pb.UpdateMoneyReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.UpdateMoneyReq;
  return proto.pb.UpdateMoneyReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.UpdateMoneyReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.UpdateMoneyReq}
 */
proto.pb.UpdateMoneyReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.UpdateMoneyReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.UpdateMoneyReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.UpdateMoneyReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.UpdateMoneyReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.pb.UpdateMoneyResp.prototype.toObject = function(opt_includeInstance) {
  return proto.pb.UpdateMoneyResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.pb.UpdateMoneyResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.UpdateMoneyResp.toObject = function(includeInstance, msg) {
  var obj = {
    userid: jspb.Message.getFieldWithDefault(msg, 1, 0),
    money: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.pb.UpdateMoneyResp}
 */
proto.pb.UpdateMoneyResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.pb.UpdateMoneyResp;
  return proto.pb.UpdateMoneyResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.pb.UpdateMoneyResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.pb.UpdateMoneyResp}
 */
proto.pb.UpdateMoneyResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUserid(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setMoney(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.pb.UpdateMoneyResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.pb.UpdateMoneyResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.pb.UpdateMoneyResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.pb.UpdateMoneyResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserid();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getMoney();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional int64 userID = 1;
 * @return {number}
 */
proto.pb.UpdateMoneyResp.prototype.getUserid = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.pb.UpdateMoneyResp.prototype.setUserid = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional int64 money = 2;
 * @return {number}
 */
proto.pb.UpdateMoneyResp.prototype.getMoney = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.pb.UpdateMoneyResp.prototype.setMoney = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * @enum {number}
 */
proto.pb.PlayerState = {
  PLAYERLOOKON: 0,
  PLAYERSITDOWN: 1,
  PLAYERAGREE: 2,
  PLAYERPLAYING: 3,
  PLAYERPICKUP: 4,
  PLAYERCALL: 5,
  PLAYERFOLLOW: 6,
  PLAYERRAISE: 7,
  PLAYERLOOK: 8,
  PLAYERCOMPARE: 9,
  PLAYERCOMPARELOSE: 10,
  PLAYEROUTCARD: 11,
  PLAYERPASS: 12,
  PLAYERCHI: 13,
  PLAYERPONG: 14,
  PLAYERMINGGANG: 15,
  PLAYERANGANG: 16,
  PLAYERTING: 17,
  PLAYERHU: 18,
  PLAYERZIMO: 19,
  PLAYERTRUSTEE: 97,
  PLAYERGIVEUP: 98,
  PLAYERSTANDUP: 99
};

goog.object.extend(exports, proto.pb);
