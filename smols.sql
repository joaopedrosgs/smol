-- ----------------------------
-- Table structure for smols
-- ----------------------------
DROP TABLE IF EXISTS "public"."smols";
CREATE TABLE "public"."smols" (
"key" serial4 NOT NULL,
"value" varchar(255) COLLATE "default" NOT NULL
)
WITH (OIDS=FALSE)

;

-- ----------------------------
-- Alter Sequences Owned By 
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table smols
-- ----------------------------
ALTER TABLE "public"."smols" ADD PRIMARY KEY ("key");
