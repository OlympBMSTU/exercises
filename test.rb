require 'pg'


class Generator 
    def initialize() 
        @subjects = ["mathematic", "physics", "informatics"]
        @tags = ["array", "dynamics", "equations", "logics", "other", "recursion", "test"]
        @path = "a"
        @right_answer = "no"
        @conn = conn = PG.connect(:dbname => 'excercieses')

    end

    def generate_subjects() 
        @subjects.each() do |subject|
            @conn.exec("INSERT INTO subject(name) values('#{subject}')")
        end
    end
    
    def generate_tables() 
        rnd = Random.new()
        p @subjects
        # @subjects.each()  do |subject|
        #     p 'a'
        #     @conn.exec("INSERT INTO SUBJECTS(name) VALUES('#{subject}')")
        # end 
        # @tags.each() do |tag| 
        #     @conn.exec("INSERT INTO TAG(name) VALUES('#{tag}')")
        # end

        0.upto(10) do |i|
            tags_arr = "'{"
            count_tags = rnd.rand(1..@tags.size)
            0.upto(count_tags) do |j|
                id = rnd.rand(0..(@tags.size() - 1))
                tags_arr += "\"" + @tags[id] + "\"" + ","
            end
            tags_arr = tags_arr.chomp(",") + "}'"
            # tags_arr += "}" 
            subject = "'" + @subjects[rnd.rand(0..(@subjects.size()-1))] + "'"
            p tags_arr
            p subject

            @conn.exec("select add_excerciese(#{rnd.rand(0..100)}, 'b', #{rnd.rand(0..5)}, 'f_', #{subject}, #{tags_arr})")
        end


        # 0.upto(20) do |i|
        #     conn.exec()
        # end

    end

end

g = Generator.new()
# g.generate_subjects()
g.generate_tables()