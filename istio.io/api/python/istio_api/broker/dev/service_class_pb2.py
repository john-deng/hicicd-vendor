# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: broker/dev/service_class.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='broker/dev/service_class.proto',
  package='istio.broker.dev',
  syntax='proto3',
  serialized_pb=_b('\n\x1e\x62roker/dev/service_class.proto\x12\x10istio.broker.dev\"o\n\x0cServiceClass\x12\x30\n\ndeployment\x18\x01 \x01(\x0b\x32\x1c.istio.broker.dev.Deployment\x12-\n\x05\x65ntry\x18\x02 \x01(\x0b\x32\x1e.istio.broker.dev.CatalogEntry\"\x1e\n\nDeployment\x12\x10\n\x08instance\x18\x01 \x01(\t\"=\n\x0c\x43\x61talogEntry\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\n\n\x02id\x18\x02 \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\tB\x19Z\x17istio.io/api/broker/devb\x06proto3')
)




_SERVICECLASS = _descriptor.Descriptor(
  name='ServiceClass',
  full_name='istio.broker.dev.ServiceClass',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='deployment', full_name='istio.broker.dev.ServiceClass.deployment', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='entry', full_name='istio.broker.dev.ServiceClass.entry', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=52,
  serialized_end=163,
)


_DEPLOYMENT = _descriptor.Descriptor(
  name='Deployment',
  full_name='istio.broker.dev.Deployment',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='instance', full_name='istio.broker.dev.Deployment.instance', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=165,
  serialized_end=195,
)


_CATALOGENTRY = _descriptor.Descriptor(
  name='CatalogEntry',
  full_name='istio.broker.dev.CatalogEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='istio.broker.dev.CatalogEntry.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='id', full_name='istio.broker.dev.CatalogEntry.id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='istio.broker.dev.CatalogEntry.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=197,
  serialized_end=258,
)

_SERVICECLASS.fields_by_name['deployment'].message_type = _DEPLOYMENT
_SERVICECLASS.fields_by_name['entry'].message_type = _CATALOGENTRY
DESCRIPTOR.message_types_by_name['ServiceClass'] = _SERVICECLASS
DESCRIPTOR.message_types_by_name['Deployment'] = _DEPLOYMENT
DESCRIPTOR.message_types_by_name['CatalogEntry'] = _CATALOGENTRY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ServiceClass = _reflection.GeneratedProtocolMessageType('ServiceClass', (_message.Message,), dict(
  DESCRIPTOR = _SERVICECLASS,
  __module__ = 'broker.dev.service_class_pb2'
  # @@protoc_insertion_point(class_scope:istio.broker.dev.ServiceClass)
  ))
_sym_db.RegisterMessage(ServiceClass)

Deployment = _reflection.GeneratedProtocolMessageType('Deployment', (_message.Message,), dict(
  DESCRIPTOR = _DEPLOYMENT,
  __module__ = 'broker.dev.service_class_pb2'
  # @@protoc_insertion_point(class_scope:istio.broker.dev.Deployment)
  ))
_sym_db.RegisterMessage(Deployment)

CatalogEntry = _reflection.GeneratedProtocolMessageType('CatalogEntry', (_message.Message,), dict(
  DESCRIPTOR = _CATALOGENTRY,
  __module__ = 'broker.dev.service_class_pb2'
  # @@protoc_insertion_point(class_scope:istio.broker.dev.CatalogEntry)
  ))
_sym_db.RegisterMessage(CatalogEntry)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z\027istio.io/api/broker/dev'))
# @@protoc_insertion_point(module_scope)
