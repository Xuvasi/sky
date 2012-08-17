#include <dbg.h>

// Initializes and compiles a query into a module.
#define COMPILE_QUERY_RAW(MODULE, ARGS, ARG_COUNT, QUERY) do {\
    struct tagbstring query = bsStatic(QUERY); \
    struct tagbstring module_name = bsStatic("eql"); \
    struct tagbstring core_class_path = bsStatic("lib/core"); \
    struct tagbstring sky_class_path = bsStatic("lib/sky"); \
    eql_compiler *compiler = eql_compiler_create(); \
    eql_compiler_add_class_path(compiler, &core_class_path); \
    eql_compiler_add_class_path(compiler, &sky_class_path); \
    int _rc = eql_compiler_compile(compiler, &module_name, &query, ARGS, ARG_COUNT, &module); \
    mu_assert(_rc == 0, "Unable to compile"); \
    if(module->error_count > 0) fprintf(stderr, "Parse error [line %d] %s\n", module->errors[0]->line_no, bdata(module->errors[0]->message)); \
    mu_assert_int_equals(module->error_count, 0); \
    eql_compiler_free(compiler); \
} while(0)

// Initializes and compiles a query into a module.
#define COMPILE_QUERY_2ARG(MODULE, ARG_TYPE1, ARG_NAME1, ARG_TYPE2, ARG_NAME2, QUERY) do {\
    uint32_t arg_count = 0; \
    eql_ast_node *args[2]; \
    if(strlen(ARG_TYPE1) > 0) { \
        struct tagbstring arg_name = bsStatic(ARG_NAME1); \
        eql_ast_node *type_ref = eql_ast_type_ref_create_cstr(ARG_TYPE1); \
        eql_ast_node *var_decl = eql_ast_var_decl_create(type_ref, &arg_name, NULL); \
        args[0] = eql_ast_farg_create(var_decl); \
        arg_count++; \
    } \
    if(strlen(ARG_TYPE2) > 0) { \
        struct tagbstring arg_name = bsStatic(ARG_NAME2); \
        eql_ast_node *type_ref = eql_ast_type_ref_create_cstr(ARG_TYPE2); \
        eql_ast_node *var_decl = eql_ast_var_decl_create(type_ref, &arg_name, NULL); \
        args[1] = eql_ast_farg_create(var_decl); \
        arg_count++; \
    } \
    COMPILE_QUERY_RAW(MODULE, args, arg_count, QUERY); \
} while(0)

#define COMPILE_QUERY_1ARG(MODULE, ARG_TYPE, ARG_NAME, QUERY) COMPILE_QUERY_2ARG(MODULE, ARG_TYPE, ARG_NAME, "", "", QUERY)
#define COMPILE_QUERY(MODULE, QUERY) COMPILE_QUERY_1ARG(MODULE, "", "", QUERY)