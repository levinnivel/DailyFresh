-- phpMyAdmin SQL Dump
-- version 5.0.3
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 12, 2022 at 04:32 PM
-- Server version: 10.4.14-MariaDB
-- PHP Version: 7.4.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `daily_fresh`
--

-- --------------------------------------------------------

--
-- Table structure for table `admin`
--

CREATE TABLE `admin` (
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `admin`
--

INSERT INTO `admin` (`user_id`) VALUES
(1);

-- --------------------------------------------------------

--
-- Table structure for table `cart`
--

CREATE TABLE `cart` (
  `id` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `cart`
--

INSERT INTO `cart` (`id`, `customer_id`) VALUES
(1, 3);

-- --------------------------------------------------------

--
-- Table structure for table `cartline`
--

CREATE TABLE `cartline` (
  `id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `goods_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `cartline`
--

INSERT INTO `cartline` (`id`, `quantity`, `goods_id`) VALUES
(1, 10, 2);

-- --------------------------------------------------------

--
-- Table structure for table `customer`
--

CREATE TABLE `customer` (
  `user_id` int(11) NOT NULL,
  `cust_address` varchar(255) NOT NULL,
  `balance` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `customer`
--

INSERT INTO `customer` (`user_id`, `cust_address`, `balance`) VALUES
(3, 'Jl. Gempol Asri', 500),
(5, 'Jl Kairo 24', 0);

-- --------------------------------------------------------

--
-- Table structure for table `goods`
--

CREATE TABLE `goods` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `price` int(11) NOT NULL,
  `description` varchar(255) NOT NULL,
  `category` varchar(255) NOT NULL,
  `stock` int(11) NOT NULL,
  `image` varchar(255) NOT NULL,
  `seller_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `goods`
--

INSERT INTO `goods` (`id`, `name`, `price`, `description`, `category`, `stock`, `image`, `seller_id`) VALUES
(1, 'Wortel 800 gr', 8000, 'Wortel segar', 'Veggies', 10, 'wortel-800-gr.png', 2),
(2, 'Kol Merah Pertiwi 500 gr', 9000, 'Kol merah segar merek Pertiwi', 'Veggies', 100, 'kol_merah_pertiwi.png', 2),
(3, 'Wagyu A5 100 gr', 150000, 'Daging Wagyu A5 produksi Jepang', 'Meat', 50, 'wagyu_a5_stark.jpg', 7),
(4, 'Roasted Garlic 192 gr', 50000, 'Seasoning/bumbu dapur', 'Spices', 100, 'roasted-garlic.jpg', 2),
(5, 'Apel Hijau 1 kg', 75000, 'Apel hijau Granny Smith', 'Fruits', 500, 'apel-hijau.jpg', 2),
(6, 'Jeruk Limau 100 gr', 8000, 'Jeruk limau langsung dari kebun', 'Fruits', 80, 'jeruk-limau.jpg', 2),
(7, 'Iga Sapi 500 gr', 70000, 'Daging iga sapi cocok untuk sop', 'Meat', 120, 'daging-iga.jpg', 7),
(8, 'Tenderloin 1 kg', 195000, 'Daging sapi has dalam/tenderloin', 'Meat', 250, 'tenderloin.jpg', 7),
(9, 'Sirloin 1 kg', 180000, 'Daging sapi has luar/sirloin', 'Meat', 250, 'sirloin.jpg', 7),
(10, 'Daging Sapi Fillet 1 kg', 140000, 'Daging sapi fillet fresh segar', 'Meat', 264, 'daging-sapi-fillet.jpg', 7),
(11, 'Thyme 100 gr', 50000, 'Premium spices thyme', 'Spices', 65, 'thyme.jpg', 2),
(12, 'Melon Fresh', 42500, 'Buah melon segar 1900 gram', 'Fruits', 15, 'melon.jpg', 6),
(13, 'Melon Import Crown', 2250000, 'Buah melon Crown import Jepang 1500 gram', 'Fruits', 10, 'melon-crown.jpg', 6),
(14, 'Bawang Bombay', 8300, 'Bawang bombay 120-170 gram/harga per pcs', 'Veggies', 76, 'bawang-bombay.jpg', 6),
(15, 'Kentang 1 kg', 19900, 'Kentang Dieng 1 kg', 'Veggies', 137, 'kentang-dieng.jpg', 6),
(16, 'Ribeye Meltique Beef 200 gr', 89000, 'Ribeye Meltique Beef (AUS) 200 gr', 'Meat', 43, 'ribeye-meltique.jpg', 6),
(17, 'Bakso Sapi 25 pcs', 26500, 'Bakso sapi bungkus isi 25 pcs/650 gr', 'Meat', 97, 'bakso-sapi.jpg', 6);

-- --------------------------------------------------------

--
-- Table structure for table `orderline`
--

CREATE TABLE `orderline` (
  `id` int(11) NOT NULL,
  `goods_id` int(11) NOT NULL,
  `order_id` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `sell_price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orderline`
--

INSERT INTO `orderline` (`id`, `goods_id`, `order_id`, `quantity`, `sell_price`) VALUES
(1, 2, 1, 10, 90000),
(2, 1, 2, 10, 80000),
(3, 2, 2, 50, 450000);

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--

CREATE TABLE `orders` (
  `id` int(11) NOT NULL,
  `rating` int(11) NOT NULL,
  `date` date NOT NULL,
  `status` int(11) NOT NULL,
  `total_price` int(11) NOT NULL,
  `customer_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`id`, `rating`, `date`, `status`, `total_price`, `customer_id`) VALUES
(1, 5, '2022-05-06', 1, 90000, 3),
(2, 5, '2022-05-06', 1, 0, 5);

-- --------------------------------------------------------

--
-- Table structure for table `payment`
--

CREATE TABLE `payment` (
  `id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `method` varchar(255) NOT NULL,
  `order_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `payment`
--

INSERT INTO `payment` (`id`, `amount`, `method`, `order_id`) VALUES
(1, 90000, 'OVO', 1),
(2, 530000, 'BCA', 2);

-- --------------------------------------------------------

--
-- Table structure for table `seller`
--

CREATE TABLE `seller` (
  `user_id` int(11) NOT NULL,
  `shop_name` varchar(255) NOT NULL,
  `website_address` varchar(255) NOT NULL,
  `seller_address` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `seller`
--

INSERT INTO `seller` (`user_id`, `shop_name`, `website_address`, `seller_address`) VALUES
(2, 'Virdimart', 'Jl. Dipati Ukur Kelurahan No.26, Lebakgede, Kecamatan Coblong, Kota Bandung, Jawa Barat 40132', 'Jl. deket RS, Bandung'),
(6, 'Wayne Foods', 'Jl. Braga No.64, Braga, Kec. Sumur Bandung, Kota Bandung, Jawa Barat 40111', 'Jl Arkham 333'),
(7, 'Stark Butchery', 'Jl. Mekar Jaya, Mekarwangi, Kec. Bojongloa Kidul, Kota Bandung, Jawa Barat 40287', 'Jl New York 53');

-- --------------------------------------------------------

--
-- Table structure for table `shipment`
--

CREATE TABLE `shipment` (
  `id` int(11) NOT NULL,
  `arrival_date` date NOT NULL,
  `ship_date` date NOT NULL,
  `order_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `shipment`
--

INSERT INTO `shipment` (`id`, `arrival_date`, `ship_date`, `order_id`) VALUES
(1, '0000-00-00', '0000-00-00', 1),
(2, '0000-00-00', '0000-00-00', 2);

-- --------------------------------------------------------

--
-- Table structure for table `ticket`
--

CREATE TABLE `ticket` (
  `id` int(11) NOT NULL,
  `category` varchar(255) NOT NULL,
  `inquiry` varchar(255) NOT NULL,
  `reply` varchar(255) NOT NULL,
  `user_id` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `ticket`
--

INSERT INTO `ticket` (`id`, `category`, `inquiry`, `reply`, `user_id`) VALUES
(1, 'Shipment', 'Kenapa barang belum sampai hingga 3 tahun lebih?', 'Coba hubungi ulang.', 3),
(2, 'Shipment', 'Kenapa belum sampai sudah 3 bulan?', 'Coba refresh ulang.', 3),
(3, 'Shipment', 'Kenapa belum sampai sudah 6 bulan?', '', 3),
(4, 'Shipment', 'Kenapa belum sampai sudah 24 bulan?', '', 3),
(5, 'Shipment', 'Kenapa belum sampai sudah 26 bulan?', '', 3);

-- --------------------------------------------------------

--
-- Table structure for table `user`
--

CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `phone` int(15) NOT NULL,
  `image_path` varchar(255) NOT NULL,
  `type_person` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `user`
--

INSERT INTO `user` (`id`, `name`, `email`, `password`, `phone`, `image_path`, `type_person`, `status`) VALUES
(1, 'Levin Martina', 'levinnivel@gmail.com', 'levinnivel', 87821429, '', 'admin', 'active'),
(2, 'Andre Viridian', 'komurasaki@gmail.com', 'komurasaki', 89765345, '', 'seller', 'active'),
(3, 'Aristo Kratos', 'deimoskratos@gmail.com', 'sideimoslah', 81567987, '', 'customer', 'active'),
(5, 'Marc Spectator', 'marc@gmail.com', 'moonknight', 2147483647, '', 'customer', 'active'),
(6, 'Bruce Banner', 'bruce112@gmail.com', 'batman', 2147483647, '', 'seller', 'active'),
(7, 'Tony Stark', 'stark@gmail.com', 'ironman', 1231351661, '', 'seller', 'active');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `admin`
--
ALTER TABLE `admin`
  ADD PRIMARY KEY (`user_id`);

--
-- Indexes for table `cart`
--
ALTER TABLE `cart`
  ADD PRIMARY KEY (`id`),
  ADD KEY `customer_id` (`customer_id`);

--
-- Indexes for table `cartline`
--
ALTER TABLE `cartline`
  ADD PRIMARY KEY (`id`,`goods_id`),
  ADD KEY `fk_goods_cart` (`goods_id`),
  ADD KEY `id` (`id`);

--
-- Indexes for table `customer`
--
ALTER TABLE `customer`
  ADD PRIMARY KEY (`user_id`);

--
-- Indexes for table `goods`
--
ALTER TABLE `goods`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_seller_goods` (`seller_id`);

--
-- Indexes for table `orderline`
--
ALTER TABLE `orderline`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_goods` (`goods_id`),
  ADD KEY `fk_order_goods` (`order_id`);

--
-- Indexes for table `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_order` (`customer_id`);

--
-- Indexes for table `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_payment` (`order_id`);

--
-- Indexes for table `seller`
--
ALTER TABLE `seller`
  ADD PRIMARY KEY (`user_id`);

--
-- Indexes for table `shipment`
--
ALTER TABLE `shipment`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_shipment` (`order_id`);

--
-- Indexes for table `ticket`
--
ALTER TABLE `ticket`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_ticket` (`user_id`);

--
-- Indexes for table `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `cart`
--
ALTER TABLE `cart`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `cartline`
--
ALTER TABLE `cartline`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `goods`
--
ALTER TABLE `goods`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `orderline`
--
ALTER TABLE `orderline`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `payment`
--
ALTER TABLE `payment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `shipment`
--
ALTER TABLE `shipment`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `ticket`
--
ALTER TABLE `ticket`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `admin`
--
ALTER TABLE `admin`
  ADD CONSTRAINT `fk_admin` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `cart`
--
ALTER TABLE `cart`
  ADD CONSTRAINT `cart_ibfk_1` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`user_id`);

--
-- Constraints for table `cartline`
--
ALTER TABLE `cartline`
  ADD CONSTRAINT `cartline_ibfk_1` FOREIGN KEY (`id`) REFERENCES `cart` (`id`),
  ADD CONSTRAINT `cartline_ibfk_2` FOREIGN KEY (`goods_id`) REFERENCES `goods` (`id`);

--
-- Constraints for table `customer`
--
ALTER TABLE `customer`
  ADD CONSTRAINT `fk_customer` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `goods`
--
ALTER TABLE `goods`
  ADD CONSTRAINT `fk_seller_goods` FOREIGN KEY (`seller_id`) REFERENCES `seller` (`user_id`);

--
-- Constraints for table `orderline`
--
ALTER TABLE `orderline`
  ADD CONSTRAINT `fk_goods` FOREIGN KEY (`goods_id`) REFERENCES `goods` (`id`),
  ADD CONSTRAINT `fk_order_goods` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

--
-- Constraints for table `orders`
--
ALTER TABLE `orders`
  ADD CONSTRAINT `fk_order` FOREIGN KEY (`customer_id`) REFERENCES `customer` (`user_id`);

--
-- Constraints for table `payment`
--
ALTER TABLE `payment`
  ADD CONSTRAINT `fk_payment` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

--
-- Constraints for table `seller`
--
ALTER TABLE `seller`
  ADD CONSTRAINT `fk_seller` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);

--
-- Constraints for table `shipment`
--
ALTER TABLE `shipment`
  ADD CONSTRAINT `fk_shipment` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`);

--
-- Constraints for table `ticket`
--
ALTER TABLE `ticket`
  ADD CONSTRAINT `fk_ticket` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
