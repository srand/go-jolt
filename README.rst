Jolt 2.0
========

This is an experimental reimplementation of Jolt in Go. The main motivation is:

 - to improve performance over the existing Python implementaiton while maintaining
   speed of development
 - to introduce a domain-specific language for build tasks
 - to improve C++ compilation times by building a complete dependency graph
   that cross task boundaries


Build task definition
---------------------

Task recipes are written in a domain-specific language. Example

  .. code-block:: c++

    workspace "ws" {
    
      project "library" {
        headers = {
          "include/header.h"
        }
    
        sources = {
          "src/tutorial.proto",
          "src/source.cpp",
        }
    
        incpaths = {
          "include",
        }
    
        macros = {
          "DEBUG",
          "FOOBAR"
        }
    
        binary = "library";
    
        task "build" {
          // Include build rules for Clang
          inherit "builtin:c++/clang/library";
    
          // And generate protobufs on the fly
          inherit "builtin:c++/protobuf";
    
          // Also run cppcheck since it's fast
          inherit "builtin:c++/cppcheck";
    
          steps {
            transform headers, sources;
          }
        }
    
        task "clang-tidy" {
          // Include build rules for Clang Tidy
          inherit "builtin:c++/clang/tidy";
    
          // Skip protobufs
          rule protobuf : skip { ext = {".proto"} }
    
          steps {
            transform headers, sources;
          }
        }
      }
    }
