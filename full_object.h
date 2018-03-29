#include <vector>

namespace dlib {
  class full_object_detection
      {
      public:
          full_object_detection(
              const rectangle& rect_,
              const std::vector<point>& parts_
          ) : rect(rect_), parts(parts_) {}

          full_object_detection(){}

          explicit full_object_detection(
              const rectangle& rect_
          ) : rect(rect_) {}

          const rectangle& get_rect() const { return rect; }
          rectangle& get_rect() { return rect; }
          unsigned long num_parts() const { return parts.size(); }

          const point& part(
              unsigned long idx
          ) const
          {
              return parts[idx];
          }

          point& part(
              unsigned long idx
          )
          {
              return parts[idx];
          }

          friend void serialize (
              const full_object_detection& item,
              std::ostream& out
          )
          {
              int version = 1;
              serialize(version, out);
              serialize(item.rect, out);
              serialize(item.parts, out);
          }

          friend void deserialize (
              full_object_detection& item,
              std::istream& in
          )
          {
              int version = 0;
              deserialize(version, in);
              if (version != 1)
                  throw serialization_error("Unexpected version encountered while deserializing dlib::full_object_detection.");

              deserialize(item.rect, in);
              deserialize(item.parts, in);
          }

          bool operator==(
              const full_object_detection& rhs
          ) const
          {
              if (rect != rhs.rect)
                  return false;
              if (parts.size() != rhs.parts.size())
                  return false;
              for (size_t i = 0; i < parts.size(); ++i)
              {
                  if (parts[i] != rhs.parts[i])
                      return false;
              }
              return true;
          }

      private:
          rectangle rect;
          std::vector<point> parts;
      };

}
