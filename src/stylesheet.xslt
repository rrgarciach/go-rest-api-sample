<?xml version="1.0"?>

<xsl:stylesheet version="1.0" xmlns:xsl="http://www.w3.org/1999/XSL/Transform">
	<xsl:output method="xml" indent="yes"/>
	<xsl:param name="values"></xsl:param>
	<xsl:param name="paramSeparator">&amp;</xsl:param>
	<xsl:param name="valueSeparator">,</xsl:param>
	<xsl:param name="nameValueSeparator">=</xsl:param>

	<xsl:template match="@* | node()">
		<xsl:copy>
			<xsl:apply-templates select="@* | node()"/>
		</xsl:copy>
	</xsl:template>

	<!-- no ends-with in xpath 1.0-->
	<xsl:template match="/fetch/entity/filter/condition[starts-with(@value,'{') and substring(@value, string-length(@value)) = '}']">
		<condition>
			<xsl:variable name="placeholder" select="translate(@value, '{}', '')" />

			<xsl:attribute name="attribute">
				<xsl:value-of select="@attribute"/>
			</xsl:attribute>
			<xsl:attribute name="operator">
				<xsl:value-of select="@operator"/>
			</xsl:attribute>

			<xsl:choose>
				<xsl:when test="contains($values, $placeholder)">
					<xsl:call-template name="chooseApplyValues">
						<xsl:with-param name="paramList" select="$values" />
						<xsl:with-param name="placeholder" select="$placeholder" />
						<xsl:with-param name="multiValue" select="(@operator='in' or @operator='not-in' or @operator='between' or @operator='not-between')" />
					</xsl:call-template>
				</xsl:when>
				<xsl:otherwise>
					<xsl:attribute name="value">
						<xsl:value-of select="@value"/>
					</xsl:attribute>
				</xsl:otherwise>
			</xsl:choose>

		</condition>
	</xsl:template>


	<xsl:template name="chooseApplyValues">
		<xsl:param name="paramList" />
		<xsl:param name="placeholder" />
		<xsl:param name="multiValue" />

		<xsl:if test="string-length($paramList)">
			<xsl:variable name="paramAndValue" select="substring-before($paramList, $paramSeparator)" />

			<xsl:choose>
				<xsl:when test="string-length($paramAndValue)">
					<xsl:call-template name="tryReplacePlaceholder">
						<xsl:with-param name="placeholder" select="$placeholder" />
						<xsl:with-param name="paramAndValue" select="$paramAndValue" />
						<xsl:with-param name="multiValue" select="$multiValue" />
					</xsl:call-template>

					<!--for the remaining parameters-->
					<xsl:call-template name="chooseApplyValues">
						<xsl:with-param name="paramList" select="substring-after($paramList, $paramSeparator)" />
						<xsl:with-param name="placeholder" select="$placeholder" />
						<xsl:with-param name="multiValue" select="$multiValue" />
					</xsl:call-template>
				</xsl:when>
				<xsl:otherwise>
					<!-- the only param=value pair left -->
					<xsl:call-template name="tryReplacePlaceholder">
						<xsl:with-param name="placeholder" select="$placeholder" />
						<xsl:with-param name="paramAndValue" select="$paramList" />
						<xsl:with-param name="multiValue" select="$multiValue" />
					</xsl:call-template>
				</xsl:otherwise>
			</xsl:choose>
		</xsl:if>
	</xsl:template>

	<xsl:template name="tryReplacePlaceholder">
		<xsl:param name="placeholder" />
		<xsl:param name="paramAndValue" />
		<xsl:param name="multiValue" />

		<xsl:if test="string-length($paramAndValue)">
			<xsl:variable name="paramName" select="substring-before($paramAndValue, $nameValueSeparator)" />

			<xsl:if test="$placeholder = $paramName">
				<xsl:variable name="value" select="substring-after($paramAndValue, $nameValueSeparator)" />

				<xsl:choose>
					<xsl:when test="$multiValue">
						<xsl:call-template name="multipleValues">
							<xsl:with-param name="text" select="$value" />
						</xsl:call-template>
					</xsl:when>
					<xsl:otherwise>
						<xsl:attribute name="value">
							<xsl:value-of select="$value" />
						</xsl:attribute>
					</xsl:otherwise>
				</xsl:choose>
			</xsl:if>
		</xsl:if>
	</xsl:template>

    <xsl:template name="multipleValues">
        <xsl:param name="text"/>

        <xsl:if test="string-length($text)">
            <value>
				<xsl:variable name="currentValue" select="substring-before($text, $valueSeparator)" />
				<xsl:choose>
					<xsl:when test="string-length($currentValue)">
						<xsl:value-of select="$currentValue"/>
					</xsl:when>
					<xsl:otherwise>
						<xsl:value-of select="$text"/>
					</xsl:otherwise>
				</xsl:choose>
            </value>

            <xsl:call-template name="multipleValues">
                <xsl:with-param name="text" select="substring-after($text, $valueSeparator)"/>
            </xsl:call-template>
        </xsl:if>
    </xsl:template>

</xsl:stylesheet>
