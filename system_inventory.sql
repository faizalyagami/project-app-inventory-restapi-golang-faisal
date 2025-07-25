PGDMP  
    )                }            system_inventory    16.3    16.3 A    =           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            >           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            @           1262    16915    system_inventory    DATABASE     �   CREATE DATABASE system_inventory WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'Indonesian_Indonesia.1252';
     DROP DATABASE system_inventory;
                postgres    false            �            1259    16929 
   categories    TABLE     f   CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);
    DROP TABLE public.categories;
       public         heap    postgres    false            �            1259    16928    categories_id_seq    SEQUENCE     �   CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.categories_id_seq;
       public          postgres    false    218            A           0    0    categories_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;
          public          postgres    false    217            �            1259    16917    items    TABLE     �  CREATE TABLE public.items (
    id integer NOT NULL,
    name character varying(100) NOT NULL,
    category_id integer NOT NULL,
    rack_id integer,
    warehouse_id integer,
    stock integer DEFAULT 0 NOT NULL,
    price integer DEFAULT 0 NOT NULL,
    min_stock integer DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.items;
       public         heap    postgres    false            �            1259    16916    items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.items_id_seq;
       public          postgres    false    216            B           0    0    items_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;
          public          postgres    false    215            �            1259    16943    racks    TABLE     a   CREATE TABLE public.racks (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);
    DROP TABLE public.racks;
       public         heap    postgres    false            �            1259    16942    racks_id_seq    SEQUENCE     �   CREATE SEQUENCE public.racks_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.racks_id_seq;
       public          postgres    false    222            C           0    0    racks_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.racks_id_seq OWNED BY public.racks.id;
          public          postgres    false    221            �            1259    17072 
   sale_items    TABLE     �   CREATE TABLE public.sale_items (
    id integer NOT NULL,
    sale_id bigint NOT NULL,
    item_id bigint NOT NULL,
    quantity bigint NOT NULL,
    price bigint NOT NULL
);
    DROP TABLE public.sale_items;
       public         heap    postgres    false            �            1259    17071    sale_items_id_seq    SEQUENCE     �   CREATE SEQUENCE public.sale_items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.sale_items_id_seq;
       public          postgres    false    229            D           0    0    sale_items_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.sale_items_id_seq OWNED BY public.sale_items.id;
          public          postgres    false    228            �            1259    17059    sales    TABLE     �   CREATE TABLE public.sales (
    id integer NOT NULL,
    user_id bigint NOT NULL,
    total bigint NOT NULL,
    created_at timestamp without time zone DEFAULT now()
);
    DROP TABLE public.sales;
       public         heap    postgres    false            �            1259    17058    sales_id_seq    SEQUENCE     �   CREATE SEQUENCE public.sales_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.sales_id_seq;
       public          postgres    false    227            E           0    0    sales_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.sales_id_seq OWNED BY public.sales.id;
          public          postgres    false    226            �            1259    17029    sessions    TABLE     �   CREATE TABLE public.sessions (
    token uuid NOT NULL,
    user_id integer,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);
    DROP TABLE public.sessions;
       public         heap    postgres    false            �            1259    17015    users    TABLE     �  CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    email character varying(100) NOT NULL,
    password text NOT NULL,
    role character varying(20) NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_role_check CHECK (((role)::text = ANY ((ARRAY['admin'::character varying, 'staff'::character varying, 'owner'::character varying])::text[])))
);
    DROP TABLE public.users;
       public         heap    postgres    false            �            1259    17014    users_id_seq    SEQUENCE     �   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    224            F           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    223            �            1259    16936 
   warehouses    TABLE     f   CREATE TABLE public.warehouses (
    id integer NOT NULL,
    name character varying(100) NOT NULL
);
    DROP TABLE public.warehouses;
       public         heap    postgres    false            �            1259    16935    warehouses_id_seq    SEQUENCE     �   CREATE SEQUENCE public.warehouses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.warehouses_id_seq;
       public          postgres    false    220            G           0    0    warehouses_id_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.warehouses_id_seq OWNED BY public.warehouses.id;
          public          postgres    false    219            x           2604    16932    categories id    DEFAULT     n   ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);
 <   ALTER TABLE public.categories ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    218    217    218            r           2604    16920    items id    DEFAULT     d   ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);
 7   ALTER TABLE public.items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    216    215    216            z           2604    16946    racks id    DEFAULT     d   ALTER TABLE ONLY public.racks ALTER COLUMN id SET DEFAULT nextval('public.racks_id_seq'::regclass);
 7   ALTER TABLE public.racks ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    222    221    222            �           2604    17075    sale_items id    DEFAULT     n   ALTER TABLE ONLY public.sale_items ALTER COLUMN id SET DEFAULT nextval('public.sale_items_id_seq'::regclass);
 <   ALTER TABLE public.sale_items ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    228    229    229            ~           2604    17062    sales id    DEFAULT     d   ALTER TABLE ONLY public.sales ALTER COLUMN id SET DEFAULT nextval('public.sales_id_seq'::regclass);
 7   ALTER TABLE public.sales ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    226    227    227            {           2604    17018    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    224    223    224            y           2604    16939    warehouses id    DEFAULT     n   ALTER TABLE ONLY public.warehouses ALTER COLUMN id SET DEFAULT nextval('public.warehouses_id_seq'::regclass);
 <   ALTER TABLE public.warehouses ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    220    219    220            /          0    16929 
   categories 
   TABLE DATA           .   COPY public.categories (id, name) FROM stdin;
    public          postgres    false    218   �H       -          0    16917    items 
   TABLE DATA           ~   COPY public.items (id, name, category_id, rack_id, warehouse_id, stock, price, min_stock, created_at, updated_at) FROM stdin;
    public          postgres    false    216   �H       3          0    16943    racks 
   TABLE DATA           )   COPY public.racks (id, name) FROM stdin;
    public          postgres    false    222   �J       :          0    17072 
   sale_items 
   TABLE DATA           K   COPY public.sale_items (id, sale_id, item_id, quantity, price) FROM stdin;
    public          postgres    false    229   �J       8          0    17059    sales 
   TABLE DATA           ?   COPY public.sales (id, user_id, total, created_at) FROM stdin;
    public          postgres    false    227   �J       6          0    17029    sessions 
   TABLE DATA           >   COPY public.sessions (token, user_id, created_at) FROM stdin;
    public          postgres    false    225   +K       5          0    17015    users 
   TABLE DATA           P   COPY public.users (id, username, email, password, role, created_at) FROM stdin;
    public          postgres    false    224   HK       1          0    16936 
   warehouses 
   TABLE DATA           .   COPY public.warehouses (id, name) FROM stdin;
    public          postgres    false    220   �L       H           0    0    categories_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.categories_id_seq', 4, true);
          public          postgres    false    217            I           0    0    items_id_seq    SEQUENCE SET     ;   SELECT pg_catalog.setval('public.items_id_seq', 32, true);
          public          postgres    false    215            J           0    0    racks_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.racks_id_seq', 4, true);
          public          postgres    false    221            K           0    0    sale_items_id_seq    SEQUENCE SET     ?   SELECT pg_catalog.setval('public.sale_items_id_seq', 4, true);
          public          postgres    false    228            L           0    0    sales_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.sales_id_seq', 3, true);
          public          postgres    false    226            M           0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 5, true);
          public          postgres    false    223            N           0    0    warehouses_id_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.warehouses_id_seq', 1, false);
          public          postgres    false    219            �           2606    16934    categories categories_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.categories DROP CONSTRAINT categories_pkey;
       public            postgres    false    218            �           2606    16927    items items_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.items DROP CONSTRAINT items_pkey;
       public            postgres    false    216            �           2606    16948    racks racks_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.racks
    ADD CONSTRAINT racks_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.racks DROP CONSTRAINT racks_pkey;
       public            postgres    false    222            �           2606    17077    sale_items sale_items_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.sale_items
    ADD CONSTRAINT sale_items_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.sale_items DROP CONSTRAINT sale_items_pkey;
       public            postgres    false    229            �           2606    17065    sales sales_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.sales
    ADD CONSTRAINT sales_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.sales DROP CONSTRAINT sales_pkey;
       public            postgres    false    227            �           2606    17034    sessions sessions_pkey 
   CONSTRAINT     W   ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (token);
 @   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_pkey;
       public            postgres    false    225            �           2606    17028    users users_email_key 
   CONSTRAINT     Q   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);
 ?   ALTER TABLE ONLY public.users DROP CONSTRAINT users_email_key;
       public            postgres    false    224            �           2606    17024    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    224            �           2606    17026    users users_username_key 
   CONSTRAINT     W   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);
 B   ALTER TABLE ONLY public.users DROP CONSTRAINT users_username_key;
       public            postgres    false    224            �           2606    16941    warehouses warehouses_pkey 
   CONSTRAINT     X   ALTER TABLE ONLY public.warehouses
    ADD CONSTRAINT warehouses_pkey PRIMARY KEY (id);
 D   ALTER TABLE ONLY public.warehouses DROP CONSTRAINT warehouses_pkey;
       public            postgres    false    220            �           2606    16979    items fk_category    FK CONSTRAINT     �   ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE CASCADE;
 ;   ALTER TABLE ONLY public.items DROP CONSTRAINT fk_category;
       public          postgres    false    218    216    4741            �           2606    16999    items fk_rack    FK CONSTRAINT        ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_rack FOREIGN KEY (rack_id) REFERENCES public.racks(id) ON DELETE SET NULL;
 7   ALTER TABLE ONLY public.items DROP CONSTRAINT fk_rack;
       public          postgres    false    222    4745    216            �           2606    17009    items fk_warehouse    FK CONSTRAINT     �   ALTER TABLE ONLY public.items
    ADD CONSTRAINT fk_warehouse FOREIGN KEY (warehouse_id) REFERENCES public.warehouses(id) ON DELETE SET NULL;
 <   ALTER TABLE ONLY public.items DROP CONSTRAINT fk_warehouse;
       public          postgres    false    4743    216    220            �           2606    17083 "   sale_items sale_items_item_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.sale_items
    ADD CONSTRAINT sale_items_item_id_fkey FOREIGN KEY (item_id) REFERENCES public.items(id) ON DELETE RESTRICT;
 L   ALTER TABLE ONLY public.sale_items DROP CONSTRAINT sale_items_item_id_fkey;
       public          postgres    false    229    4739    216            �           2606    17078 "   sale_items sale_items_sale_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.sale_items
    ADD CONSTRAINT sale_items_sale_id_fkey FOREIGN KEY (sale_id) REFERENCES public.sales(id) ON DELETE CASCADE;
 L   ALTER TABLE ONLY public.sale_items DROP CONSTRAINT sale_items_sale_id_fkey;
       public          postgres    false    227    229    4755            �           2606    17066    sales sales_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.sales
    ADD CONSTRAINT sales_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 B   ALTER TABLE ONLY public.sales DROP CONSTRAINT sales_user_id_fkey;
       public          postgres    false    224    4749    227            �           2606    17035    sessions sessions_user_id_fkey    FK CONSTRAINT     �   ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;
 H   ALTER TABLE ONLY public.sessions DROP CONSTRAINT sessions_user_id_fkey;
       public          postgres    false    4749    224    225            /   E   x�3�t�I�.)�����2�H-J�I,I�S�N�+�/�2�t�.N-�/�,V���-(-I-�2�&���� 09T      -   u  x����j�0��ӧ8/�$'�����TV�c��3hPi*����C7FKC.|�����,��E(�%�����c�@�C:B3��1N7�7=�A��L���F��+	*�X9_�
�w1x�I!�lðe��#ϴȔ�c�M�K)AM�p��:T0��	���_�QL�����nK�?��&���G�c����5������~[6�ߖ�����Ŝ4�$�u'�!+ltn���`���V�&�����;���"���B޺DQ��Q��א�;x�P�s�gyʢ��`�m�K7���-J�8�B�Al�M�Ԗ�h�s{�Gx�`,�G��6�����c���eV�֙ �Pu������k�u6TE>༵��e�vཟ$�'����      3   $   x�3�J�Vp4�23����gC.#F��� �a�      :   $   x�3�4�4�4�42 . ��$di ����� ]s      8   2   x�3�4�4571 N##S]s]3#+#K+C=sKC�=... �z�      6      x������ � �      5   z  x���ˎ�@�5<�Y������nVJ�o/�1���j�@��炓L<�YͮR��_�T� >�^S���� U�F��� X�8�jj�ꇝ�f�/��3o�Gb�n8@�r�:�Hhw�]�X�l*7��D�F1eO�z��6�l"��0\74T�������a�u��qY�b��G����B��8�:k�e�>x=��ԝw��Ƽ�{�b3�8g�1���K�*�k�=�+�8�7?�ivo�Q���y��'�@ޮ'{�6C\X�+�ݭ/;X�\�ʗq�����oo�m��#�3uV]� �Gu�ʏ%�A������&��7�N5�Qk��d{�2$�Mn������?Տ΄"�j�H��w�I�@      1   :   x�3�t/MI�KW-I�M�2�q����1��Z���X������[PZ�Z����� ��     