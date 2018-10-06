require 'pg'


class Generator 
    def initialize() 
        @subjects = ["mathematic", "physics", "informatics"]
        @tags = ["array", "dynamics", "equations", "logics", "other", "recursion", "test"]
        @path = "a"
        @right_answer = "no"
        @conn = conn = PG.connect(:dbname => 'excercieses')

    end
    
    def generate_tables() 
        p @subjects
        # @subjects.each()  do |subject|
        #     p 'a'
        #     @conn.exec("INSERT INTO SUBJECTS(name) VALUES('#{subject}')")
        # end 
        # @tags.each() do |tag| 
        #     @conn.exec("INSERT INTO TAG(name) VALUES('#{tag}')")
        # end

        0.upto(10000) do |i|
            @conn.exec("insert into exceriese(author_id, right_answer, level, file_name, subject) VALUES(3, #{i}, #{i}, '242', 'math')")
            @conn.exec("insert into tag_excerciese(tag_id, excerciese_id) values(4, #{i+2})")
        end


        # 0.upto(20) do |i|
        #     conn.exec()
        # end

    end

end

# def main() 
    g = Generator.new()
    g.generate_tables()
# end