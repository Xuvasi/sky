#include <stdio.h>
#include <stdlib.h>

#include <add_property_message.h>
#include <mem.h>

#include "../minunit.h"


//==============================================================================
//
// Test Cases
//
//==============================================================================

//--------------------------------------
// Serialization
//--------------------------------------

int test_sky_add_property_message_pack() {
    cleantmp();
    sky_add_property_message *message = sky_add_property_message_create();
    message->property->type = SKY_PROPERTY_TYPE_OBJECT;
    message->property->data_type = SKY_DATA_TYPE_INT;
    message->property->name = bfromcstr("foo");
    
    FILE *file = fopen("tmp/message", "w");
    mu_assert_bool(sky_add_property_message_pack(message, file) == 0);
    fclose(file);
    mu_assert_file("tmp/message", "tests/fixtures/add_property_message/0/message");
    sky_add_property_message_free(message);
    return 0;
}

int test_sky_add_property_message_unpack() {
    FILE *file = fopen("tests/fixtures/add_property_message/0/message", "r");
    sky_add_property_message *message = sky_add_property_message_create();
    mu_assert_bool(sky_add_property_message_unpack(message, file) == 0);
    fclose(file);

    mu_assert_bool(message->property->type == SKY_PROPERTY_TYPE_OBJECT);
    mu_assert_bool(message->property->data_type == SKY_DATA_TYPE_INT);
    mu_assert_bstring(message->property->name, "foo");
    sky_add_property_message_free(message);
    return 0;
}


//--------------------------------------
// Processing
//--------------------------------------

int test_sky_add_property_message_process() {
    cleantmp();
    sky_table *table = sky_table_create();
    table->path = bfromcstr("tmp");
    sky_table_open(table);
    
    sky_add_property_message *message = sky_add_property_message_create();
    message->property->type = SKY_PROPERTY_TYPE_OBJECT;
    message->property->data_type = SKY_DATA_TYPE_INT;
    message->property->name = bfromcstr("foo");

    FILE *output = fopen("tmp/output", "w");
    mu_assert(sky_add_property_message_process(message, table, output) == 0, "");
    fclose(output);
    mu_assert_file("tmp/properties", "tests/fixtures/add_property_message/1/table/properties");
    mu_assert_file("tmp/output", "tests/fixtures/add_property_message/1/output");

    sky_add_property_message_free(message);
    sky_table_free(table);
    return 0;
}


//==============================================================================
//
// Setup
//
//==============================================================================

int all_tests() {
    mu_run_test(test_sky_add_property_message_pack);
    mu_run_test(test_sky_add_property_message_unpack);
    mu_run_test(test_sky_add_property_message_process);
    return 0;
}

RUN_TESTS()