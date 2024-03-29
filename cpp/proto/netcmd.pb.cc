// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: netcmd.proto

#include "netcmd.pb.h"

#include <algorithm>

#include <google/protobuf/io/coded_stream.h>
#include <google/protobuf/extension_set.h>
#include <google/protobuf/wire_format_lite.h>
#include <google/protobuf/descriptor.h>
#include <google/protobuf/generated_message_reflection.h>
#include <google/protobuf/reflection_ops.h>
#include <google/protobuf/wire_format.h>
// @@protoc_insertion_point(includes)
#include <google/protobuf/port_def.inc>

PROTOBUF_PRAGMA_INIT_SEG
namespace netcmd {
constexpr Billboard::Billboard(
  ::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized)
  : msg_(&::PROTOBUF_NAMESPACE_ID::internal::fixed_address_empty_string){}
struct BillboardDefaultTypeInternal {
  constexpr BillboardDefaultTypeInternal()
    : _instance(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized{}) {}
  ~BillboardDefaultTypeInternal() {}
  union {
    Billboard _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT BillboardDefaultTypeInternal _Billboard_default_instance_;
constexpr EnterGsNotify::EnterGsNotify(
  ::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized)
  : servertime_(&::PROTOBUF_NAMESPACE_ID::internal::fixed_address_empty_string)
  , serveropentime_(&::PROTOBUF_NAMESPACE_ID::internal::fixed_address_empty_string)
  , serverversion_(0u)
  , clientfuncswitch_(0u){}
struct EnterGsNotifyDefaultTypeInternal {
  constexpr EnterGsNotifyDefaultTypeInternal()
    : _instance(::PROTOBUF_NAMESPACE_ID::internal::ConstantInitialized{}) {}
  ~EnterGsNotifyDefaultTypeInternal() {}
  union {
    EnterGsNotify _instance;
  };
};
PROTOBUF_ATTRIBUTE_NO_DESTROY PROTOBUF_CONSTINIT EnterGsNotifyDefaultTypeInternal _EnterGsNotify_default_instance_;
}  // namespace netcmd
static ::PROTOBUF_NAMESPACE_ID::Metadata file_level_metadata_netcmd_2eproto[2];
static constexpr ::PROTOBUF_NAMESPACE_ID::EnumDescriptor const** file_level_enum_descriptors_netcmd_2eproto = nullptr;
static constexpr ::PROTOBUF_NAMESPACE_ID::ServiceDescriptor const** file_level_service_descriptors_netcmd_2eproto = nullptr;

const ::PROTOBUF_NAMESPACE_ID::uint32 TableStruct_netcmd_2eproto::offsets[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::netcmd::Billboard, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::netcmd::Billboard, msg_),
  ~0u,  // no _has_bits_
  PROTOBUF_FIELD_OFFSET(::netcmd::EnterGsNotify, _internal_metadata_),
  ~0u,  // no _extensions_
  ~0u,  // no _oneof_case_
  ~0u,  // no _weak_field_map_
  PROTOBUF_FIELD_OFFSET(::netcmd::EnterGsNotify, servertime_),
  PROTOBUF_FIELD_OFFSET(::netcmd::EnterGsNotify, serverversion_),
  PROTOBUF_FIELD_OFFSET(::netcmd::EnterGsNotify, serveropentime_),
  PROTOBUF_FIELD_OFFSET(::netcmd::EnterGsNotify, clientfuncswitch_),
};
static const ::PROTOBUF_NAMESPACE_ID::internal::MigrationSchema schemas[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) = {
  { 0, -1, sizeof(::netcmd::Billboard)},
  { 6, -1, sizeof(::netcmd::EnterGsNotify)},
};

static ::PROTOBUF_NAMESPACE_ID::Message const * const file_default_instances[] = {
  reinterpret_cast<const ::PROTOBUF_NAMESPACE_ID::Message*>(&::netcmd::_Billboard_default_instance_),
  reinterpret_cast<const ::PROTOBUF_NAMESPACE_ID::Message*>(&::netcmd::_EnterGsNotify_default_instance_),
};

const char descriptor_table_protodef_netcmd_2eproto[] PROTOBUF_SECTION_VARIABLE(protodesc_cold) =
  "\n\014netcmd.proto\022\006netcmd\"\030\n\tBillboard\022\013\n\003m"
  "sg\030\001 \001(\t\"l\n\rEnterGsNotify\022\022\n\nserverTime\030"
  "\001 \001(\t\022\025\n\rserverVersion\030\002 \001(\r\022\026\n\016serverOp"
  "enTime\030\003 \001(\t\022\030\n\020clientfuncswitch\030\004 \001(\rb\006"
  "proto3"
  ;
static ::PROTOBUF_NAMESPACE_ID::internal::once_flag descriptor_table_netcmd_2eproto_once;
const ::PROTOBUF_NAMESPACE_ID::internal::DescriptorTable descriptor_table_netcmd_2eproto = {
  false, false, 166, descriptor_table_protodef_netcmd_2eproto, "netcmd.proto", 
  &descriptor_table_netcmd_2eproto_once, nullptr, 0, 2,
  schemas, file_default_instances, TableStruct_netcmd_2eproto::offsets,
  file_level_metadata_netcmd_2eproto, file_level_enum_descriptors_netcmd_2eproto, file_level_service_descriptors_netcmd_2eproto,
};
PROTOBUF_ATTRIBUTE_WEAK ::PROTOBUF_NAMESPACE_ID::Metadata
descriptor_table_netcmd_2eproto_metadata_getter(int index) {
  ::PROTOBUF_NAMESPACE_ID::internal::AssignDescriptors(&descriptor_table_netcmd_2eproto);
  return descriptor_table_netcmd_2eproto.file_level_metadata[index];
}

// Force running AddDescriptors() at dynamic initialization time.
PROTOBUF_ATTRIBUTE_INIT_PRIORITY static ::PROTOBUF_NAMESPACE_ID::internal::AddDescriptorsRunner dynamic_init_dummy_netcmd_2eproto(&descriptor_table_netcmd_2eproto);
namespace netcmd {

// ===================================================================

class Billboard::_Internal {
 public:
};

Billboard::Billboard(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor();
  RegisterArenaDtor(arena);
  // @@protoc_insertion_point(arena_constructor:netcmd.Billboard)
}
Billboard::Billboard(const Billboard& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  msg_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  if (!from._internal_msg().empty()) {
    msg_.Set(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, from._internal_msg(), 
      GetArena());
  }
  // @@protoc_insertion_point(copy_constructor:netcmd.Billboard)
}

void Billboard::SharedCtor() {
msg_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}

Billboard::~Billboard() {
  // @@protoc_insertion_point(destructor:netcmd.Billboard)
  SharedDtor();
  _internal_metadata_.Delete<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

void Billboard::SharedDtor() {
  GOOGLE_DCHECK(GetArena() == nullptr);
  msg_.DestroyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}

void Billboard::ArenaDtor(void* object) {
  Billboard* _this = reinterpret_cast< Billboard* >(object);
  (void)_this;
}
void Billboard::RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena*) {
}
void Billboard::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}

void Billboard::Clear() {
// @@protoc_insertion_point(message_clear_start:netcmd.Billboard)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  msg_.ClearToEmpty();
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* Billboard::_InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::PROTOBUF_NAMESPACE_ID::uint32 tag;
    ptr = ::PROTOBUF_NAMESPACE_ID::internal::ReadTag(ptr, &tag);
    CHK_(ptr);
    switch (tag >> 3) {
      // string msg = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 10)) {
          auto str = _internal_mutable_msg();
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(::PROTOBUF_NAMESPACE_ID::internal::VerifyUTF8(str, "netcmd.Billboard.msg"));
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      default: {
      handle_unusual:
        if ((tag & 7) == 4 || tag == 0) {
          ctx->SetLastTag(tag);
          goto success;
        }
        ptr = UnknownFieldParse(tag,
            _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
            ptr, ctx);
        CHK_(ptr != nullptr);
        continue;
      }
    }  // switch
  }  // while
success:
  return ptr;
failure:
  ptr = nullptr;
  goto success;
#undef CHK_
}

::PROTOBUF_NAMESPACE_ID::uint8* Billboard::_InternalSerialize(
    ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:netcmd.Billboard)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string msg = 1;
  if (this->msg().size() > 0) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_msg().data(), static_cast<int>(this->_internal_msg().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "netcmd.Billboard.msg");
    target = stream->WriteStringMaybeAliased(
        1, this->_internal_msg(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:netcmd.Billboard)
  return target;
}

size_t Billboard::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:netcmd.Billboard)
  size_t total_size = 0;

  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string msg = 1;
  if (this->msg().size() > 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_msg());
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    return ::PROTOBUF_NAMESPACE_ID::internal::ComputeUnknownFieldsSize(
        _internal_metadata_, total_size, &_cached_size_);
  }
  int cached_size = ::PROTOBUF_NAMESPACE_ID::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void Billboard::MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:netcmd.Billboard)
  GOOGLE_DCHECK_NE(&from, this);
  const Billboard* source =
      ::PROTOBUF_NAMESPACE_ID::DynamicCastToGenerated<Billboard>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:netcmd.Billboard)
    ::PROTOBUF_NAMESPACE_ID::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:netcmd.Billboard)
    MergeFrom(*source);
  }
}

void Billboard::MergeFrom(const Billboard& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:netcmd.Billboard)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.msg().size() > 0) {
    _internal_set_msg(from._internal_msg());
  }
}

void Billboard::CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:netcmd.Billboard)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void Billboard::CopyFrom(const Billboard& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:netcmd.Billboard)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool Billboard::IsInitialized() const {
  return true;
}

void Billboard::InternalSwap(Billboard* other) {
  using std::swap;
  _internal_metadata_.Swap<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(&other->_internal_metadata_);
  msg_.Swap(&other->msg_, &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
}

::PROTOBUF_NAMESPACE_ID::Metadata Billboard::GetMetadata() const {
  return GetMetadataStatic();
}


// ===================================================================

class EnterGsNotify::_Internal {
 public:
};

EnterGsNotify::EnterGsNotify(::PROTOBUF_NAMESPACE_ID::Arena* arena)
  : ::PROTOBUF_NAMESPACE_ID::Message(arena) {
  SharedCtor();
  RegisterArenaDtor(arena);
  // @@protoc_insertion_point(arena_constructor:netcmd.EnterGsNotify)
}
EnterGsNotify::EnterGsNotify(const EnterGsNotify& from)
  : ::PROTOBUF_NAMESPACE_ID::Message() {
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  servertime_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  if (!from._internal_servertime().empty()) {
    servertime_.Set(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, from._internal_servertime(), 
      GetArena());
  }
  serveropentime_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  if (!from._internal_serveropentime().empty()) {
    serveropentime_.Set(::PROTOBUF_NAMESPACE_ID::internal::ArenaStringPtr::EmptyDefault{}, from._internal_serveropentime(), 
      GetArena());
  }
  ::memcpy(&serverversion_, &from.serverversion_,
    static_cast<size_t>(reinterpret_cast<char*>(&clientfuncswitch_) -
    reinterpret_cast<char*>(&serverversion_)) + sizeof(clientfuncswitch_));
  // @@protoc_insertion_point(copy_constructor:netcmd.EnterGsNotify)
}

void EnterGsNotify::SharedCtor() {
servertime_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
serveropentime_.UnsafeSetDefault(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
::memset(reinterpret_cast<char*>(this) + static_cast<size_t>(
    reinterpret_cast<char*>(&serverversion_) - reinterpret_cast<char*>(this)),
    0, static_cast<size_t>(reinterpret_cast<char*>(&clientfuncswitch_) -
    reinterpret_cast<char*>(&serverversion_)) + sizeof(clientfuncswitch_));
}

EnterGsNotify::~EnterGsNotify() {
  // @@protoc_insertion_point(destructor:netcmd.EnterGsNotify)
  SharedDtor();
  _internal_metadata_.Delete<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

void EnterGsNotify::SharedDtor() {
  GOOGLE_DCHECK(GetArena() == nullptr);
  servertime_.DestroyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
  serveropentime_.DestroyNoArena(&::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited());
}

void EnterGsNotify::ArenaDtor(void* object) {
  EnterGsNotify* _this = reinterpret_cast< EnterGsNotify* >(object);
  (void)_this;
}
void EnterGsNotify::RegisterArenaDtor(::PROTOBUF_NAMESPACE_ID::Arena*) {
}
void EnterGsNotify::SetCachedSize(int size) const {
  _cached_size_.Set(size);
}

void EnterGsNotify::Clear() {
// @@protoc_insertion_point(message_clear_start:netcmd.EnterGsNotify)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  servertime_.ClearToEmpty();
  serveropentime_.ClearToEmpty();
  ::memset(&serverversion_, 0, static_cast<size_t>(
      reinterpret_cast<char*>(&clientfuncswitch_) -
      reinterpret_cast<char*>(&serverversion_)) + sizeof(clientfuncswitch_));
  _internal_metadata_.Clear<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>();
}

const char* EnterGsNotify::_InternalParse(const char* ptr, ::PROTOBUF_NAMESPACE_ID::internal::ParseContext* ctx) {
#define CHK_(x) if (PROTOBUF_PREDICT_FALSE(!(x))) goto failure
  while (!ctx->Done(&ptr)) {
    ::PROTOBUF_NAMESPACE_ID::uint32 tag;
    ptr = ::PROTOBUF_NAMESPACE_ID::internal::ReadTag(ptr, &tag);
    CHK_(ptr);
    switch (tag >> 3) {
      // string serverTime = 1;
      case 1:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 10)) {
          auto str = _internal_mutable_servertime();
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(::PROTOBUF_NAMESPACE_ID::internal::VerifyUTF8(str, "netcmd.EnterGsNotify.serverTime"));
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      // uint32 serverVersion = 2;
      case 2:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 16)) {
          serverversion_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      // string serverOpenTime = 3;
      case 3:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 26)) {
          auto str = _internal_mutable_serveropentime();
          ptr = ::PROTOBUF_NAMESPACE_ID::internal::InlineGreedyStringParser(str, ptr, ctx);
          CHK_(::PROTOBUF_NAMESPACE_ID::internal::VerifyUTF8(str, "netcmd.EnterGsNotify.serverOpenTime"));
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      // uint32 clientfuncswitch = 4;
      case 4:
        if (PROTOBUF_PREDICT_TRUE(static_cast<::PROTOBUF_NAMESPACE_ID::uint8>(tag) == 32)) {
          clientfuncswitch_ = ::PROTOBUF_NAMESPACE_ID::internal::ReadVarint32(&ptr);
          CHK_(ptr);
        } else goto handle_unusual;
        continue;
      default: {
      handle_unusual:
        if ((tag & 7) == 4 || tag == 0) {
          ctx->SetLastTag(tag);
          goto success;
        }
        ptr = UnknownFieldParse(tag,
            _internal_metadata_.mutable_unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(),
            ptr, ctx);
        CHK_(ptr != nullptr);
        continue;
      }
    }  // switch
  }  // while
success:
  return ptr;
failure:
  ptr = nullptr;
  goto success;
#undef CHK_
}

::PROTOBUF_NAMESPACE_ID::uint8* EnterGsNotify::_InternalSerialize(
    ::PROTOBUF_NAMESPACE_ID::uint8* target, ::PROTOBUF_NAMESPACE_ID::io::EpsCopyOutputStream* stream) const {
  // @@protoc_insertion_point(serialize_to_array_start:netcmd.EnterGsNotify)
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  // string serverTime = 1;
  if (this->servertime().size() > 0) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_servertime().data(), static_cast<int>(this->_internal_servertime().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "netcmd.EnterGsNotify.serverTime");
    target = stream->WriteStringMaybeAliased(
        1, this->_internal_servertime(), target);
  }

  // uint32 serverVersion = 2;
  if (this->serverversion() != 0) {
    target = stream->EnsureSpace(target);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::WriteUInt32ToArray(2, this->_internal_serverversion(), target);
  }

  // string serverOpenTime = 3;
  if (this->serveropentime().size() > 0) {
    ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::VerifyUtf8String(
      this->_internal_serveropentime().data(), static_cast<int>(this->_internal_serveropentime().length()),
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::SERIALIZE,
      "netcmd.EnterGsNotify.serverOpenTime");
    target = stream->WriteStringMaybeAliased(
        3, this->_internal_serveropentime(), target);
  }

  // uint32 clientfuncswitch = 4;
  if (this->clientfuncswitch() != 0) {
    target = stream->EnsureSpace(target);
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::WriteUInt32ToArray(4, this->_internal_clientfuncswitch(), target);
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    target = ::PROTOBUF_NAMESPACE_ID::internal::WireFormat::InternalSerializeUnknownFieldsToArray(
        _internal_metadata_.unknown_fields<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(::PROTOBUF_NAMESPACE_ID::UnknownFieldSet::default_instance), target, stream);
  }
  // @@protoc_insertion_point(serialize_to_array_end:netcmd.EnterGsNotify)
  return target;
}

size_t EnterGsNotify::ByteSizeLong() const {
// @@protoc_insertion_point(message_byte_size_start:netcmd.EnterGsNotify)
  size_t total_size = 0;

  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  // Prevent compiler warnings about cached_has_bits being unused
  (void) cached_has_bits;

  // string serverTime = 1;
  if (this->servertime().size() > 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_servertime());
  }

  // string serverOpenTime = 3;
  if (this->serveropentime().size() > 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::StringSize(
        this->_internal_serveropentime());
  }

  // uint32 serverVersion = 2;
  if (this->serverversion() != 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::UInt32Size(
        this->_internal_serverversion());
  }

  // uint32 clientfuncswitch = 4;
  if (this->clientfuncswitch() != 0) {
    total_size += 1 +
      ::PROTOBUF_NAMESPACE_ID::internal::WireFormatLite::UInt32Size(
        this->_internal_clientfuncswitch());
  }

  if (PROTOBUF_PREDICT_FALSE(_internal_metadata_.have_unknown_fields())) {
    return ::PROTOBUF_NAMESPACE_ID::internal::ComputeUnknownFieldsSize(
        _internal_metadata_, total_size, &_cached_size_);
  }
  int cached_size = ::PROTOBUF_NAMESPACE_ID::internal::ToCachedSize(total_size);
  SetCachedSize(cached_size);
  return total_size;
}

void EnterGsNotify::MergeFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_merge_from_start:netcmd.EnterGsNotify)
  GOOGLE_DCHECK_NE(&from, this);
  const EnterGsNotify* source =
      ::PROTOBUF_NAMESPACE_ID::DynamicCastToGenerated<EnterGsNotify>(
          &from);
  if (source == nullptr) {
  // @@protoc_insertion_point(generalized_merge_from_cast_fail:netcmd.EnterGsNotify)
    ::PROTOBUF_NAMESPACE_ID::internal::ReflectionOps::Merge(from, this);
  } else {
  // @@protoc_insertion_point(generalized_merge_from_cast_success:netcmd.EnterGsNotify)
    MergeFrom(*source);
  }
}

void EnterGsNotify::MergeFrom(const EnterGsNotify& from) {
// @@protoc_insertion_point(class_specific_merge_from_start:netcmd.EnterGsNotify)
  GOOGLE_DCHECK_NE(&from, this);
  _internal_metadata_.MergeFrom<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(from._internal_metadata_);
  ::PROTOBUF_NAMESPACE_ID::uint32 cached_has_bits = 0;
  (void) cached_has_bits;

  if (from.servertime().size() > 0) {
    _internal_set_servertime(from._internal_servertime());
  }
  if (from.serveropentime().size() > 0) {
    _internal_set_serveropentime(from._internal_serveropentime());
  }
  if (from.serverversion() != 0) {
    _internal_set_serverversion(from._internal_serverversion());
  }
  if (from.clientfuncswitch() != 0) {
    _internal_set_clientfuncswitch(from._internal_clientfuncswitch());
  }
}

void EnterGsNotify::CopyFrom(const ::PROTOBUF_NAMESPACE_ID::Message& from) {
// @@protoc_insertion_point(generalized_copy_from_start:netcmd.EnterGsNotify)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

void EnterGsNotify::CopyFrom(const EnterGsNotify& from) {
// @@protoc_insertion_point(class_specific_copy_from_start:netcmd.EnterGsNotify)
  if (&from == this) return;
  Clear();
  MergeFrom(from);
}

bool EnterGsNotify::IsInitialized() const {
  return true;
}

void EnterGsNotify::InternalSwap(EnterGsNotify* other) {
  using std::swap;
  _internal_metadata_.Swap<::PROTOBUF_NAMESPACE_ID::UnknownFieldSet>(&other->_internal_metadata_);
  servertime_.Swap(&other->servertime_, &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  serveropentime_.Swap(&other->serveropentime_, &::PROTOBUF_NAMESPACE_ID::internal::GetEmptyStringAlreadyInited(), GetArena());
  ::PROTOBUF_NAMESPACE_ID::internal::memswap<
      PROTOBUF_FIELD_OFFSET(EnterGsNotify, clientfuncswitch_)
      + sizeof(EnterGsNotify::clientfuncswitch_)
      - PROTOBUF_FIELD_OFFSET(EnterGsNotify, serverversion_)>(
          reinterpret_cast<char*>(&serverversion_),
          reinterpret_cast<char*>(&other->serverversion_));
}

::PROTOBUF_NAMESPACE_ID::Metadata EnterGsNotify::GetMetadata() const {
  return GetMetadataStatic();
}


// @@protoc_insertion_point(namespace_scope)
}  // namespace netcmd
PROTOBUF_NAMESPACE_OPEN
template<> PROTOBUF_NOINLINE ::netcmd::Billboard* Arena::CreateMaybeMessage< ::netcmd::Billboard >(Arena* arena) {
  return Arena::CreateMessageInternal< ::netcmd::Billboard >(arena);
}
template<> PROTOBUF_NOINLINE ::netcmd::EnterGsNotify* Arena::CreateMaybeMessage< ::netcmd::EnterGsNotify >(Arena* arena) {
  return Arena::CreateMessageInternal< ::netcmd::EnterGsNotify >(arena);
}
PROTOBUF_NAMESPACE_CLOSE

// @@protoc_insertion_point(global_scope)
#include <google/protobuf/port_undef.inc>
